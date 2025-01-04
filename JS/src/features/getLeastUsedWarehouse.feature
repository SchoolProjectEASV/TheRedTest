Feature: GetLeastUsedWarehouse - Boundary Tests

  #------------------------------------------
  # Scenario 1: Same-day range
  #------------------------------------------
  Scenario: Single day usage
    Given I have 2 warehouses with following data:
      | ID | totalVolume | usageOn2025-01-10 |
      | 1  | 100.0       | 40.0             |
      | 2  | 100.0       | 20.0             |
    When I call GetLeastUsedWarehouse from "2025-01-10" to "2025-01-10"
    Then the least used warehouse should be "2"

  #------------------------------------------
  # Scenario 2: Invalid date range
  #------------------------------------------
  Scenario: Start date after end date
    Given I have 1 warehouse
    When I call GetLeastUsedWarehouse from "2025-01-12" to "2025-01-11"
    Then an error should be returned with message "the start date cannot be later than the end date"

  #------------------------------------------
  # Scenario 3: No warehouses
  #------------------------------------------
  Scenario: Empty warehouse list
    Given I have 0 warehouses
    When I call GetLeastUsedWarehouse from "2025-01-10" to "2025-01-11"
    Then an error should be returned with message "No warehouses available"

  #------------------------------------------
  # Scenario 4: All usage = 0
  #------------------------------------------
  Scenario: Zero usage
    Given I have 2 warehouses and the following data:
      | ID | totalVolume | usageOn2025-01-10 |
      | 1  | 100.0       | 0.0              |
      | 2  | 100.0       | 0.0              |
    When I call GetLeastUsedWarehouse from "2025-01-10" to "2025-01-10"
    Then the result should be -1
    And no error is returned
