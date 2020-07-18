package app

var MainPage = `
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>BarCode scanner</title>
    <style>
        input, label {
            display: block;
        }
    </style>
</head>
<body class="text-center">
<form action="/code" method="POST">
    <p>
        <label for="code">BarCode (Tap Enter)</label>
        <input id="code" type="text" name="code" autofocus/>
    </p>
</form>

<form action="/code" method="POST">
    <p>
        <label for="c_code">BarCode only</label>
        <input id="c_code" type="text" name="c_code"/>
    </p>
</form>

<script>
    window.onload = function () {
        document.getElementById("shk_code").focus();
    };
</script>
</body>
</html>
`
