# GoInjection
![GoInjection Banner](img.png)

<p align="center">
  <a href="#Features">🔧 Features</a> | 
  <a href="#WAFIdentification">🛡️ WAF Identification</a> | 
  <a href="#Fingerprinting">🔍 Fingerprinting</a> | 
  <a href="#InjectionTypes">💥 Injection Types</a> | 
  <a href="#GUI">🎨 GUI</a> | 
  <a href="#SQLQueryBuilder">⚙️ SQL Query Builder</a>
</p>

<p align="center">
  <a href="https://github.com/Axion-Security/GoInjection/actions">
    <img alt="CI Status" src="https://img.shields.io/github/actions/workflow/status/Axion-Security/GoInjection/go.yml?branch=main">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License" src="https://img.shields.io/badge/License-CC NonCommercial-blue.svg">
  </a>
  <a href="https://github.com/yourusername/GoInjection/stargazers">
    <img alt="Stars" src="https://img.shields.io/github/stars/Axion-Security/GoInjection">
  </a>
</p>

## 🚀 Features

### 🛡️ WAF Identification
- Detect Web Application Firewalls (WAFs) by analyzing HTTP headers and server responses.
- Identifies potential WAFs based on changes in response status codes and content patterns.

### 🔍 Fingerprinting
- Automatically detects the type of database by analyzing errors and executing advanced SQL injection techniques like UNION and stacked queries.
- Custom tool-generated queries ensure precise fingerprinting.

### ⚙️ Interpreter
- Automatically selects the correct SQL dialect (MySQL, PostgreSQL, MSSQL, etc.) based on the target DBMS.
- Ensures the right syntax for effective SQL injections.

### 🧩 Resolver
- In development, this feature helps identify table columns and the name of the current database for more advanced injections.
- Essential for gaining deeper access to database structures.

### 💥 Injection Types
- Supports **Blind Injections** (Boolean and Time-based), **Error-based Injections**, and **UNION-based Injections**.
- Queries for these injection types are automatically created by the tool’s custom query generator.

### 🎨 GUI (Graphical User Interface)
- Modern, user-friendly interface designed to simplify the process, making it accessible for both experienced pentesters and beginners.
- Features intuitive controls for both new and experienced users.

### ⚙️ SQL Query Builder
- Generates tailored SQL payloads for each DBMS, optimizing query structure, including custom openings, endings, and elements.
- Fully automated query builder for each injection type.

## 📦 Installation

### Prerequisites:
Ensure you have Python 3.x installed and the required dependencies:

```bash
pip install flask pymysql
```

### Setup:
Clone the repository and install dependencies:

```bash
git clone https://github.com/Axion-Security/GoInjection.git
pip install flask pymysql
```

## 📝 Usage Example

Set up a basic SQL Injection Lab for testing:

```sql
USE sql_injection_lab;

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

INSERT INTO products (name, price) VALUES
('Laptop', 799.99),
('Smartphone', 599.99),
('Tablet', 399.99),
('Headphones', 49.99),
('Smartwatch', 199.99);
```

## 🔑 License

This project is licensed under the [License](/LICENSE). See the `LICENSE` file for more details.