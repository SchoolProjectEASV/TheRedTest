Feature: CalculateAvailableCapacity - Boundary Tests

  #------------------------------------------
  # Scenario 1: Same start/end date
  #------------------------------------------
  Scenario: Single-day range
    Given I have "1" warehouse with total volume "100.0"
    And warehouse usage on "2025-01-10" is "20.0"
    When I call CalculateAvailableCapacity from "2025-01-10" to "2025-01-10"
    Then the available capacities should be:
      | 2025-01-10 | 80.0 |

  #------------------------------------------
  # Scenario 2: Invalid date range
  #------------------------------------------
  Scenario: Start date after end date
    Given I have "1" warehouse
    When I call CalculateAvailableCapacity from "2025-01-12" to "2025-01-11"
    Then an error should be returned with message "the start date cannot be later than the end date"

  #------------------------------------------
  # Scenario 3: No warehouses
  #------------------------------------------
  Scenario: Empty warehouse list
    Given I have "0" warehouses
    When I call CalculateAvailableCapacity from "2025-01-10" to "2025-01-11"
    Then an error should be returned with message "no warehouses available"

  #------------------------------------------
  # Scenario 4: Fully booked
  #------------------------------------------
  Scenario: 100% usage
    Given I have "1" warehouse with total volume "100.0"
    And warehouse usage on "2025-01-10" is "100.0"
    When I call CalculateAvailableCapacity from "2025-01-10" to "2025-01-10"
    Then the available capacities should be:
      | 2025-01-10 | 0.0 |
