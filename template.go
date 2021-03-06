package main

const indexTemplate string = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Redirector</title>

        <!-- Latest compiled and minified CSS -->
        <link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css">
        <link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.min.css">

        <!-- Optional theme -->
        <style type="text/css" media="all">
            /* Space out content a bit */
            body {
                padding-top: 20px;
                padding-bottom: 20px;
            }

            /* Everything but the jumbotron gets side spacing for mobile first views */
            .header,
            .marketing,
            .footer {
                padding-right: 15px;
                padding-left: 15px;
            }

            /* Custom page header */
            .header {
                border-bottom: 1px solid #e5e5e5;
            }
            /* Make the masthead heading the same height as the navigation */
            .header h3 {
                padding-bottom: 19px;
                margin-top: 0;
                margin-bottom: 0;
                line-height: 40px;
            }

            /* Custom page footer */
            .footer {
                padding-top: 19px;
                color: #777;
                border-top: 1px solid #e5e5e5;
            }

            /* Customize container */
            @media (min-width: 768px) {
                .container {
                    max-width: 730px;
                }
            }
            .container-narrow > hr {
                margin: 30px 0;
            }

            /* Main marketing message and sign up button */
            .jumbotron {
                text-align: center;
                border-bottom: 1px solid #e5e5e5;
            }
            .jumbotron .btn {
                padding: 14px 24px;
                font-size: 21px;
            }

            /* Supporting marketing content */
            .marketing {
                margin: 40px 0;
            }
            .marketing p + h4 {
                margin-top: 28px;
            }

            /* Responsive: Portrait tablets and up */
            @media screen and (min-width: 768px) {
                /* Remove the padding we set earlier */
                .header,
                .marketing,
                .footer {
                    padding-right: 0;
                    padding-left: 0;
                }
                /* Space out the masthead */
                .header {
                    margin-bottom: 30px;
                }
                /* Remove the bottom border on the jumbotron for visual effect */
                .jumbotron {
                    border-bottom: 0;
                }
            }

        </style>

        <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
        <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
        <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
        <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
        <![endif]-->
    </head>
    <body>
        <div class="container">
            <div class="header">
                <ul class="nav nav-pills pull-right">
                    <li class="active"><a href="#">Home</a></li>
                    <li><a href="https://github.com/pboehm/redirector">
                        <i class="fa fa-github fa-lg"></i> Code</a></li>
                </ul>
                <h3 class="text-muted">Redirector</h3>
            </div>

            <div class="jumbotron">
                <h1>Redirect all the things ...</h1>

                <p class="lead">Redirector makes your own domain even more
                useful, by redirecting custom subdomains to any location
                on the web.</p>

                <hr />

                <form class="form-inline" role="form" method="GET"
                      action="/admin/create">

                    <div id="redirector_input" class="form-group">
                        <input id="fqdn" name="fqdn" class="form-control input-lg"
                               type="text" placeholder="g.example.org">

                        <i style="margin-left: 20px; margin-right: 20px;"
                           class="fa fa-arrow-right fa-lg"></i>

                       <input id="target" name="target" class="form-control input-lg"
                              type="text" placeholder="github.com/pboehm">
                    </div>

                    <hr />

                    <input type="submit" id="register"
                        class="btn btn-primary disabled" value="Redirect it!" />
                </form>
            </div>

            <div id="hint" class="alert alert-info" role="alert"
                 style="display: none;">

                <strong>Attention!</strong> You have to create a CNAME record <code id="cname_fqdn">testttt</code> pointing to <code>{{ .fqdn }}</code>
            </div>

            <div class="footer">
                <p>&copy; Philipp Böhm</p>
            </div>

        </div> <!-- /container -->

        <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->

        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
        <!-- Latest compiled and minified JavaScript -->
        <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>

        <script type="text/javascript" charset="utf-8">

            function isValid() {
                $('#register').removeClass("disabled");
                $('#redirector_input').removeClass("has-error");
                $('#redirector_input').addClass("has-success");
                $('#hint').show();
                $('#cname_fqdn').text($('#fqdn').val());
            }

            function isNotValid(argument) {
                $('#register').addClass("disabled");
                $('#redirector_input').removeClass("has-success");
                $('#redirector_input').addClass("has-error");
                $('#hint').hide();
            }

            function validate() {
                var hostname = $('#fqdn').val();
                var target = $('#target').val();

                if( hostname == "" || target == "") {
                    isNotValid();
                } else {
                    $.getJSON("/admin/available/" + hostname + "/", function( data ) {
                        if (data.available) {
                            isValid();
                        } else {
                            isNotValid();
                        }
                    }).error(function(){ isNotValid(); });
                }
            }

            $(document).ready(function() {
                var timer = null;
                $('#redirector_input input').on('keydown', function () {
                    clearTimeout(timer);
                    timer = setTimeout(validate, 800)
                });
            });
        </script>
    </body>
</html>
`
