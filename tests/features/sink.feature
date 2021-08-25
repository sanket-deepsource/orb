Feature: Sink management



    Scenario Outline: Create sink with "<name_validation>" sink name, "<desc_validation>" description, "<tags_validation>" tags, "<type_validation>" type, "<rh_validation>" remote host and "<usrnm_validation>" username
        Given User is logged in ORB site (email: "<email>", password: "<password>")
             And is on the sink management page
        When Click on "+ NEW SINK" button
            And Enter name label ("<name_label>")
            And Enter descripton ("<description>")
            And Choose sink type
            And Click next button
            And Enter remote host ("<remote_host>")
            And Enter username ("<username>")
            And Enter password ("<sink_password>")
            And Click next button
            And Click save button
        Then A new sink ("<result>") be created


        Examples:
            | result | name_validation | desc_validation | tags_validation | type_validation | rh_validation | usrnm_validation|      email     | password | name_label |    description   | remote_host |   username  |  sink_password |
            | should |      valid      |       valid     |     valid       |       valid     |     valid     |     valid       | test@email.com | 12345678 |  test_sink | test_create_sink |  localhost  | user_tester | 963852741 |
        #     |      valid      |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |       any       |        any      |      any        |        any      |       any     |      any        |
        #     |       any       |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |    invalid      |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |        any      |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |    invalid      |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |       valid     |      any      |     valid       |
        #     |      valid      |       valid     |     valid       |       valid     |    invalid    |     valid       |


    Scenario Outline: Close delete sink with name label
        Given User is logged in ORB site (email: "<email>", password: "<password>")
             And is on the sink management page
             And a sink already exist ("<name_label>")
        When Choose "delete sink" option
            # And Enter the name of the sink to be deleted ("<name_label>")
            # And Press 'CLOSE' button

        Examples:
            |      email     | password | name_label |
            | test@email.com | 12345678 | test_sink  |
        
    Scenario Outline: Close delete sink without name label
        Given User is logged in ORB site (email: "<email>", password: "<password>")
             And is on the sink management page
             And a sink already exist ("<name_label>")
        When Choose "delete sink" option
            # And Enter the name of the sink to be deleted ("<name_label>")
            # And Press 'CLOSE' button

        Examples:
            |      email     | password | name_label |
            | test@email.com | 12345678 | test_sink  |


    Scenario Outline: Visualize an existing sink
        Given User is logged in ORB site (email: "<email>", password: "<password>")
            And is on the sink management page
            And a sink already exist ("<name_label>")
        When Press visualize sink button
        Then Sink details page should be displayed

        Examples:
            |      email     | password | name_label |
            | test@email.com | 12345678 | test_sink  |

    Scenario Outline: Delete sink
        Given User is logged in ORB site (email: "<email>", password: "<password>")
            And is on the sink management page
            And a sink already exist ("<name_label>")
        When Choose "delete sink" option
            And Enter the name of the sink to be deleted ("<name_label>")
            And Press delete button
        Then Referred sink ("<name_label>") should be deleted

        Examples:
            |      email     | password | name_label |
            | test@email.com | 12345678 | test_sink  |
