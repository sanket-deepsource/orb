Feature: ORB login


    Scenario: Open register an account page
        Given User is on the ORB login page
        When Press register link 
        Then Register an account page should be displayed
    
    
    Scenario Outline: Register an account
        Given User is on the register page
        When Complete the form with "<fullname>" full name ("<name>"), "<email>" email address ("<email_adr>"), "<pwd>" password ("<pwd1>") and "<rep_pwd>" password on repeat ("<pwd2>")
        Then Register button "<mode>" be enabled
            And New account "<mode>" be registered
            And Login using referred email address and password "<mode>" be enabled

        Examples:
            | fullname |  email  |   pwd   | rep_pwd   | name |   email_adr    |   pwd1   |    pwd2    |    mode   |
            |   valid  |  valid  |  valid  |    same   | orb-test  | test@email.com | 12345678 | 12345678   |   should  |
            |   valid  |  valid  |  valid  | different | orb-test  | test1@email.com | 12345678 | 123456780  | shouldn't |


    Scenario Outline: Login using "<type1>" email address and "<type2>" password
        Given User is on the ORB login page
        When Enter "<type1>" email address ("<email>")
            And Enter "<type2>" password ("<password>")
            And Press the ORB button Log In
        Then Access to ORB "<mode>" be enabled


        Examples:
            |     type1    |   type2   |    mode   |     email       | password  |
            | unregistered |  correct  | shouldn't | test2@email.com | 12345678  |
            | registered   | incorrect | shouldn't | test1@email.com | 123456780 |
            | unregistered | incorrect | shouldn't | test2@email.com | 123456780 |
            | registered   |  correct  |   should  | test1@email.com | 12345678  |
