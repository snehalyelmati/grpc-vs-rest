from flask import Flask, json, jsonify, request
# from waitress import serve

app = Flask(__name__)

@app.route('/hello', methods=['POST'])
def sayHello():
    message = request.get_json()['Name']
    return message

if __name__ == "__main__":
    app.run(host='127.0.0.1', port=5052)
    # serve(app, host='127.0.0.1', port=5052)
