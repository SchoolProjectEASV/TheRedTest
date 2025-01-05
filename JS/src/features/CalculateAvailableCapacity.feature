Feature: CalculateAvailableCapacity

Scenario: Single-day range
  Given I have 1 warehouse with total volume 100.0
  
  And warehouse usage is:
    | date       | usage |
    | 2025-01-10 | 20.0  |
  When I call CalculateAvailableCapacity from "2025-01-10" to "2025-01-10"
  Then the available capacities should be:
    | date       | capacity |
    | 2025-01-10 | 80.0     |

Scenario: Start date after end date
  Given I have 1 default warehouse

  When I call CalculateAvailableCapacity from "2025-01-12" to "2025-01-11"
  Then an error should be returned with message "the start date cannot be later than the end date"

Scenario: Empty warehouse list
  Given I have 0 warehouses
  When I call CalculateAvailableCapacity from "2025-01-10" to "2025-01-11"
  Then an error should be returned with message "no warehouses available"

Scenario: 100% usage
  Given I have 1 warehouse with total volume 100.0
  And warehouse usage is:
    | date       | usage |
    | 2025-01-10 | 100.0 |
  When I call CalculateAvailableCapacity from "2025-01-10" to "2025-01-10"
  Then the available capacities should be:
    | date       | capacity |
    | 2025-01-10 | 0.0      |
