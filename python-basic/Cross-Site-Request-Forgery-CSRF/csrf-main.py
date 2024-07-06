# from flask import Flask, render_template, request, redirect, url_for, flash
# from flask_wtf.csrf import CSRFProtect

# app = Flask(__name__)
# app.secret_key = 'your_secret_key'  # Necessary for CSRF protection
# csrf = CSRFProtect(app)

# @app.route('/transfer', methods=['POST'])
# def transfer():
#     try:
#         amount = float(request.form['amount'])
#         destination_account = request.form['destination_account']
#         # Logic to transfer funds...
#         # Assuming transfer logic is implemented here.
        
#         flash('Transfer successful!', 'success')
#     except ValueError:
#         flash('Invalid amount. Please enter a valid number.', 'danger')
#     except Exception as e:
#         flash(f'An error occurred: {e}', 'danger')

#     return redirect(url_for('dashboard'))

# @app.route('/dashboard')
# def dashboard():
#     return render_template('dashboard.html')

# if __name__ == '__main__':
#     app.run()


# from flask import Flask, render_template, request, redirect, url_for, flash, jsonify
# from flask_wtf.csrf import CSRFProtect, generate_csrf

# app = Flask(__name__)
# app.secret_key = 'your_secret_key'  # Necessary for CSRF protection
# csrf = CSRFProtect(app)

# @app.route('/transfer', methods=['POST'])
# def transfer():
#     try:
#         amount = float(request.form['amount'])
#         destination_account = request.form['destination_account']
#         # Logic to transfer funds...
#         # Assuming transfer logic is implemented here.
        
#         flash('Transfer successful!', 'success')
#     except ValueError:
#         flash('Invalid amount. Please enter a valid number.', 'danger')
#     except Exception as e:
#         flash(f'An error occurred: {e}', 'danger')

#     return redirect(url_for('dashboard'))

# @app.route('/dashboard')
# def dashboard():
#     return render_template('dashboard.html')

# @app.route('/get_csrf_token', methods=['GET'])
# def get_csrf_token():
#     return jsonify({'csrf_token': generate_csrf()})

# if __name__ == '__main__':
#     app.run()

# from flask import Flask, render_template, request, redirect, url_for, flash, jsonify
# from flask_wtf.csrf import CSRFProtect, generate_csrf

# app = Flask(__name__)
# app.secret_key = 'your_secret_key'  # Necessary for CSRF protection
# csrf = CSRFProtect(app)

# @app.route('/transfer', methods=['POST'])
# def transfer():
#     try:
#         amount = float(request.form['amount'])
#         destination_account = request.form['destination_account']
#         # Logic to transfer funds...
#         # Assuming transfer logic is implemented here.
        
#         flash('Transfer successful!', 'success')
#     except ValueError:
#         flash('Invalid amount. Please enter a valid number.', 'danger')
#     except Exception as e:
#         flash(f'An error occurred: {e}', 'danger')

#     return redirect(url_for('dashboard'))

# @app.route('/dashboard')
# def dashboard():
#     return render_template('dashboard.html')

# @app.route('/get_csrf_token', methods=['GET'])
# def get_csrf_token():
#     return jsonify({'csrf_token': generate_csrf()})

# if __name__ == '__main__':
#     app.run()

from flask import Flask, render_template, request, redirect, url_for, flash, jsonify, session
from flask_wtf.csrf import CSRFProtect, generate_csrf
from flask_session import Session

app = Flask(__name__)
app.secret_key = 'your_secret_key'  # Necessary for CSRF protection and session management

# Session configuration
app.config['SESSION_TYPE'] = 'filesystem'
Session(app)
csrf = CSRFProtect(app)

@app.route('/transfer', methods=['POST'])
def transfer():
    try:
        amount = float(request.form['amount'])
        destination_account = request.form['destination_account']
        # Logic to transfer funds...
        # Assuming transfer logic is implemented here.
        
        flash('Transfer successful!', 'success')
    except ValueError:
        flash('Invalid amount. Please enter a valid number.', 'danger')
    except Exception as e:
        flash(f'An error occurred: {e}', 'danger')

    return redirect(url_for('dashboard'))

@app.route('/dashboard')
def dashboard():
    return render_template('dashboard.html')

@app.route('/get_csrf_token', methods=['GET'])
def get_csrf_token():
    csrf_token = generate_csrf()
    session['csrf_token'] = csrf_token
    return jsonify({'csrf_token': csrf_token})

if __name__ == '__main__':
    app.run()
