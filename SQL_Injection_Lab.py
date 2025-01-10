from flask import Flask, request, render_template_string
import pymysql

app = Flask(__name__)

def get_db_connection():
    """Establish and return a database connection using PyMySQL."""
    return pymysql.connect(
        host="localhost",
        user="root",
        password="",
        database="sql_injection_lab",
        cursorclass=pymysql.cursors.DictCursor
    )

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

@app.route('/product', methods=['GET'])
def product():
    product_id = request.args.get('id')

    conn = get_db_connection()
    cursor = conn.cursor()

    # Vulnerable query (using string formatting directly)
    query = "SELECT * FROM products WHERE id = '%s'" % product_id  # SQL Injection vulnerability
    try:
        cursor.execute(query)
        product = cursor.fetchone()
        conn.close()

        if product:
            return f"<h2>Product Details</h2><p>ID: {product['id']}</p><p>Name: {product['name']}</p><p>Price: {product['price']}</p>"
        else:
            return "<h2>No product found</h2>"
    except pymysql.Error as err:
        conn.close()
        return f"<h2>SQL Error:</h2><pre>{err}</pre>"
    except Exception as e:
        conn.close()
        return f"<h2>Unexpected Error:</h2><pre>{str(e)}</pre>"

if __name__ == '__main__':
    app.run(debug=True)
