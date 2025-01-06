# TheRedTest

## Overview
This project demonstrates the use of Behavior-Driven Development (BDD). The system handles key warehouse functionalities such as capacity calculation, least-used warehouse identification, and ensuring item compatibility with warehouse dimensions.

## Application Features
The warehouse management system includes the following methods:

1. **FindAvailableWarehouse**: Identifies available warehouses based on a date range and item dimensions.
2. **GetFullyUtilizedDates**: Determines dates when warehouses are fully utilized.
3. **CalculateAvailableCapacity**: Computes available storage capacities for specific date ranges.
4. **GetLeastUsedWarehouse**: Finds the warehouse with the lowest utilization over a given date range.

This system focuses on ensuring efficient warehouse operations while validating critical inputs such as item dimensions, date ranges, and capacity constraints.

## Framework used

**GoDog (GoLang)**

**Cucumber.js (JavaScript)**

**XUnit.Gherkin.Quick (.NET)**

## Running the Tests

### **GoDog**
1. Navigate to the project directory, and then navigate to the "godogs" directory.
2. Run the tests using: `go test -v`.

### **Cucumber.js**
1. Navigate to the project directory.
2. Execute the tests using: `npm test`.

### **XUnit.Gherkin.Quick**
1. Open the project in Visual Studio.
2. Use the built-in test runner or .NET CLI to execute the tests.
