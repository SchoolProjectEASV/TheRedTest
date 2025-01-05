Feature: GetFullyUtilizedDates

  #------------------------------------------
  # Scenario 1: Same start/end date
  #------------------------------------------
  Scenario: Single-day range
    Given I have 1 warehouse with total volume 100.0
    And warehouse usage on "2025-01-10" is 50.0
    When I call GetFullyUtilizedDates from "2025-01-10" to "2025-01-10"
    Then the fully utilized dates should be:
      | date |

  #------------------------------------------
  # Scenario 2: Invalid date range
  #------------------------------------------
  Scenario: Start date after end date
    Given I have 1 warehouse with total volume 100.0
    When I call GetFullyUtilizedDates from "2025-01-12" to "2025-01-11"
    Then an error should be returned with message "the start date cannot be later than the end date"

  #------------------------------------------
  # Scenario 3: No warehouses
  #------------------------------------------
  Scenario: Empty warehouse list
    Given I have 0 warehouses
    When I call GetFullyUtilizedDates from "2025-01-10" to "2025-01-11"
    Then an error should be returned with message "no warehouses available"

  #------------------------------------------
  # Scenario 4: 100% usage on exactly one day
  #------------------------------------------
  Scenario: Exactly one fully utilized day
    Given I have 1 warehouse with total volume 100.0
    And warehouse usage on "2025-01-10" is 100.0
    And warehouse usage on "2025-01-11" is 50.0
    When I call GetFullyUtilizedDates from "2025-01-10" to "2025-01-11"
    Then the fully utilized dates should be:
      | date       |
      | 2025-01-10 |