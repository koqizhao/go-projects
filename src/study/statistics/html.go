package main

const _HTML_TOP = `
<!DOCTYPE html>
<html>
    <head>
        <title>Number Statistics</title>
        <style type="text/css">
            * {
                margin: 0;
                padding: 0;
                font-family: "微软雅黑", Consola, Arial
            }
            #main {
                margin: 0 auto;
                padding: 20px;
            }
            p {
                margin: 20px 0;
                font-size: 14px;
            }
            p.error {
                display: none;
            }
            p.error.show {
                display: block;
            }
            table {
                margin-top: 30px;
                border-collapse: collapse;
                border: 1px solid #000000;
            }
            table caption {
                font-weight: bold;
            }
            th, td {
                text-align: left;
                border: 1px solid #000000;
                padding: 5px 10px;
                font-size: 14px;
            }
            th {
                width: 80px;
            }
            td {
                width: 420px;
            }
            input[type=text] {
                display: block;
                margin: 5px 0;
                width: 400px;
            }
        </style>
    </head>
    <body>
        <div id="main">
            <h3>Statistics</h3>
            <p>Computes basic statistics for a given list of numbers</p>
`

const _HTML_BOTTOM = `
        </div>
    </body>
</html>
`
const _HTML_FORM = `
            <form method="post">
                <p>
                    <label for="numbers">Numbers (comma or space separated)</label>
                    <input name="numbers" type="text" value="%s" />
                    <input name="caculate" type="submit" value="Calculate" />
                </p>
                <p class="error %s">
                    <span class="title">Error:</span>
                    <span>%s</span>
                <p>
            </form>
`
const _HTML_RESULT_TABLE = `
            <table>
                <caption>Results</caption>
                <tr>
                    <th>Numbers</th>
                    <td>[ %s ]</td>
                </tr>
                <tr>
                    <th>Count</th>
                    <td>%d</td>
                </tr>
                <tr>
                    <th>Mean</th>
                    <td>%f</td>
                </tr>
                <tr>
                    <th>Median</th>
                    <td>%f</td>
                </tr>
            </table>
`
