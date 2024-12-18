from flask import Flask, request, render_template_string
import mysql.connector
import re

app = Flask(__name__)

def get_db_connection():
    return mysql.connector.connect(
        host="localhost",
        user="root",
        password="",
        database="sql_injection_lab"
    )

def waf_filter(input_data):
    patterns = [
        r"('|\-\-|\;|--|\bUNION\b|\bSELECT\b|\bINSERT\b|\bDROP\b|\bUPDATE\b)",
        r"(\bOR\s1=1\b)",
    ]

    for pattern in patterns:
        if re.search(pattern, input_data, re.IGNORECASE):
            return False

    return True

@app.route('/')
def index():
    return render_template_string('''
    <h1>SQL Injection Lab - Products</h1>
    <form action="/product" method="get">
        <label for="id">Product ID:</label>
        <input type="text" id="id" name="id">
        <br><br>
        <input type="submit" value="View Product">
    </form>
    ''')

@app.route('/waf_product', methods=['GET'])
def waf_product():
    product_id = request.args.get('id')

    if not waf_filter(product_id):
        return "<h2>Potential SQL Injection attempt blocked by WAF</h2>"

    conn = get_db_connection()
    cursor = conn.cursor()

    query = f"SELECT * FROM products WHERE id = '{product_id}'"
    try:
        cursor.execute(query)
        product = cursor.fetchone()
        conn.close()
        if product:
            return f"<h2>Product Details</h2><p>ID: {product[0]}</p><p>Name: {product[1]}</p><p>Price: {product[2]}</p>"
        else:
            return "<h2>No product found</h2>"
    except mysql.connector.Error as err:
        conn.close()
        return f"<h2>SQL Error: {err}</h2>"
    except Exception as e:
        conn.close()
        return f"<h2>Error: {str(e)}</h2>"

@app.route('/product', methods=['GET'])
def product():
    product_id = request.args.get('id')

    conn = get_db_connection()
    cursor = conn.cursor()

    query = f"SELECT * FROM products WHERE id = '{product_id}'"
    try:
        cursor.execute(query)
        product = cursor.fetchone()
        conn.close()
        if product:
            return f"<h2>Product Details</h2><p>ID: {product[0]}</p><p>Name: {product[1]}</p><p>Price: {product[2]}</p>"
        else:
            return "<h2>No product found</h2>"
    except mysql.connector.Error as err:
        conn.close()
        return f"<h2>SQL Error: {err}</h2>"
    except Exception as e:
        conn.close()
        return f"<h2>Error: {str(e)}</h2>"

if __name__ == '__main__':
    app.run(debug=True)
