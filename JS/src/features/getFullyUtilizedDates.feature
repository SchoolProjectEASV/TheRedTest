Feature: GetFullyUtilizedDates - Boundary Tests

  #------------------------------------------
  # Scenario 1: Single-day range
  #------------------------------------------
  Scenario: Single-day range
    Given I have 1 warehouse with the given volume "100.0"
    And warehouse usage on "2025-01-10" is 50.0
    When I call GetFullyUtilizedDates from "2025-01-10" to "2025-01-10"
    Then the fully utilized dates should be "[]"
    # "[]" if usage (50.0) < total (100.0), or ["2025-01-10"] if usage == total

  #------------------------------------------
  # Scenario 2: Invalid date range
  #------------------------------------------
  Scenario: Start date after end date
    Given I have 1 warehouse with the complete volume 100.0
    When I call GetFullyUtilizedDates from "2025-01-12" to "2025-01-11"
    Then an error should be returned with message "the start date cannot be later than the end date"

  #------------------------------------------
  # Scenario 3: Empty warehouse list
  #------------------------------------------
  Scenario: No warehouses
    Given I have 0 warehouses
    When I call GetFullyUtilizedDates from "2025-01-10" to "2025-01-11"
    Then an error should be returned with message "no warehouses available"

  #------------------------------------------
  # Scenario 4: Exactly one fully utilized day
  #------------------------------------------
  Scenario: Exactly one fully utilized day
    Given I have 1 warehouse with total volume 100.0
    And warehouse usage on "2025-01-10" is 100.0
    And warehouse usage on "2025-01-11" is 50.0
    When I call GetFullyUtilizedDates from "2025-01-10" to "2025-01-11"
    Then the fully utilized dates should be "[\"2025-01-10\"]"
