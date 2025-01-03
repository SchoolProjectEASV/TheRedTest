Feature: FindAvailableWarehouse

  #------------------------------------------
  # Scenario 1: Same day start/end
  #------------------------------------------
  Scenario: Start Date == End Date
    Given today is "2025-01-09"
    And I have 1 warehouse with total volume 1.0
    And the warehouse usage is empty on all days
    When I call FindAvailableWarehouse from "2025-01-10" to "2025-01-10" with dimensions:
      | height | width | length |
      | 1.0    | 1.0   | 1.0    |
    Then I should receive warehouse ID 1

  #------------------------------------------
  # Scenario 2: Start date after end date
  #------------------------------------------
  Scenario: Invalid date range (start after end)
    Given today is "2025-01-09"
    And I have 1 warehouse with total volume 1.0
    And the warehouse usage is empty on all days
    When I call FindAvailableWarehouse from "2025-01-10" to "2025-01-09" with dimensions:
      | height | width | length |
      | 1.0    | 1.0   | 1.0    |
    Then an error should be returned with message "start date cannot be later than end date"

  #------------------------------------------
  # Scenario 3: Zero dimension
  #------------------------------------------
  Scenario: Zero or negative dimension
    Given today is "2025-01-09"
    And I have 1 warehouse with total volume 10.0
    When I call FindAvailableWarehouse from "2025-01-10" to "2025-01-11" with dimensions:
      | height | width | length |
      | 0.0    | 1.0   | 1.0    |
    Then an error should be returned with message "the 3d model has invalid dimensions (zero or negative)"

  #------------------------------------------
  # Scenario 4: Exactly matching warehouse capacity
  #------------------------------------------
  Scenario: Required volume == Warehouse capacity
    Given today is "2025-01-09"
    And I have 1 warehouse with total volume 10.0
    And the warehouse usage is empty on all days
    When I call FindAvailableWarehouse from "2025-01-10" to "2025-01-11" with dimensions:
      | height | width | length |
      | 2.15   | 2.15  | 2.15   |
    Then I should receive warehouse ID 1

  #------------------------------------------
  # Scenario 5: No warehouses
  #------------------------------------------
  Scenario: Empty warehouse list
    Given today is "2025-01-09"
    And I have 0 warehouses
    When I call FindAvailableWarehouse from "2025-01-10" to "2025-01-11" with dimensions:
      | height | width | length |
      | 1.0    | 1.0   | 1.0    |
    Then an error should be returned with message "no warehouses available"