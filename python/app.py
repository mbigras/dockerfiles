from flask import Flask

app = Flask(__name__)

@app.route('/', defaults={'u_path': ''})
@app.route('/<path:u_path>')
def handler(u_path):
    return "URL.Path = {}\n".format(u_path)
