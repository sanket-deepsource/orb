from time import sleep
from behave import given, when, then
import functions


login_page = "http://localhost:80/auth/login"
orb_page =  "http://localhost/pages/home"
login_fail_page = 'http://localhost/auth/login'
alert_message_fail_login  = 'Login/Email combination is not correct, please try again.'
register_account_page= 'http://localhost/auth/register'

#---------
@given("User is on the ORB login page")
def check_login_page(context):
    functions.go_to_page(login_page, context)

@when("Press register link")
def open_register_account_page(context):
    functions.press_button('.text-link', context)

@then("Register an account page should be displayed")
def check_register_account_page(context):
    functions.check(register_account_page, context)

#----------

@given("User is on the register page")
def check_register_page(context):
    functions.go_to_page(register_account_page, context)

@when('Complete the form with "{fullname}" full name ("{name}"), "{email}" email address ("{email_adr}"), "{pwd}" password ("{pwd1}") and "{rep_pwd}" password on repeat ("{pwd2}")')
def complete_form(context, fullname, name, email, email_adr, pwd, pwd1, rep_pwd, pwd2):
    functions.enter_information('input-name', name, context, source_by=0)
    functions.enter_information('input-email', email_adr, context, source_by=0)
    functions.enter_information('input-password', pwd1, context, source_by=0)
    functions.enter_information('input-re-password', pwd2, context, source_by=0)

@then('Register button "{mode}" be enabled')
def press_register_button(context, mode):
    if mode == 'should':
        assert functions.get_button_attribute('.appearance-filled', "aria-disabled", context) == 'false'
        functions.press_button('.appearance-filled', context)
    elif mode == "shouldn't":
        #assert functions.get_button_attribute('.appearance-filled', "aria-disabled", context) == 'true'
        assert 'btn-disabled' in functions.get_button_attribute('.appearance-filled', "class", context)
    else:
        raise ValueError('invalide mode')

@then('New account "{mode}" be registered')
def check_registered_account(context, mode):
    if mode == "should":
        functions.check(orb_page, context)
    elif mode =="shouldn't":
        functions.check(register_account_page, context)

@then('Login using referred email address and password "{mode}" be enabled')
def login_using_new_account(context, mode):
    pass


#----------
@when('Enter "{type1}" email address ("{email}")')
def enter_username(context, type1, email):
    functions.enter_information('input-email', email, context, source_by=0)

@when('Enter "{type2}" password ("{password}")')
def enter_password(context, type2, password):
    functions.enter_information('input-password', password, context, source_by=0)

@when('Press the ORB button Log In')
def click_button(context):
    functions.press_button('.appearance-filled', context)

@then('Access to ORB "{mode}" be enabled')
def check_access(context, mode):
    if mode == 'should':
        functions.check(orb_page, context)
    else:
        text = functions.get_text_from_alert('alert-message', context)
        assert text == alert_message_fail_login
        functions.check(login_fail_page, context)
