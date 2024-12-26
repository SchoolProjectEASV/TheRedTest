Feature: Warehouse Storage Service

  Scenario: Find available warehouse for item within date range
    Given the following warehouses exist:
      | Id | Height | Width | Length |
      | 1  | 10     | 10    | 10     |
      | 2  | 20     | 20    | 20     |
    And the following items exist in warehouse 1:
      | ItemId | ItemName | Height | Width | Length | StartDate           | EndDate             | IsActive |
      | 1      | Item1    | 2      | 2     | 2      | 2024-12-25T00:00:00 | 2024-12-30T23:59:59 | true     |
    When I search for an available warehouse from "2024-12-27" to "2024-12-28" ...
    Then the available warehouse ID should be 1

  # Add more scenarios to cover other methods and edge cases
