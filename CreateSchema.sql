CREATE SCHEMA AzureQl;
GO

CREATE TABLE AzureQl.Department(
    DepartmentID INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    Name VARCHAR(50)
);
GO

CREATE TABLE AzureQl.Employee(
    EmployeeID INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    FirstName VARCHAR(32),
    LastName VARCHAR(32),
    Username VARCHAR(32),
    EncryptedPassword VARCHAR(256),
    Email VARCHAR(64),
    DOB VARCHAR(10),
    DepartmentID INT REFERENCES AzureQl.Department(DepartmentID) ON UPDATE CASCADE ON DELETE SET NULL,
    Position VARCHAR(32)
);
GO

SELECT * FROM AzureQl.Employee;
GO

-- sqlcmd -S localhost:1433 -U sa -P 'superStrong(!)Password' -d master -i ./CreateSchema.sql