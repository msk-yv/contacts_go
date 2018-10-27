<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf8">
        <title>Удаление контакта </title>
        <link rel="stylesheet" href="style.css" />
    </head>
    <body>
        <?php include('views/header.html'); ?>
        <div class="wraperFirstTable">
            <table class="insert_search">
                <tr>
                    <th>Внести новый контакт</th>
                    <th>Поиск по имени</th>
                </tr> 
                <tr>
                    <td><?php include('views/insertForm.php'); ?></td>
                    <td><?php include('views/searchForm.php'); ?></td>
                </tr>
            </table>
        </div>
    
        <div class="wraperSecondTable">
            <h2>Вы уверены, что хотите удалить контакт:</h2>
            <table class="results">
                <tr>
                    <th>Имя контакта</th>
                    <th>Телефон</th>
                    <th>Электронная почта</th>
                    <th></th>
                </tr> 
                <?php foreach ($contacts as $contact) : ?>
                <form action="/models/delete.php" method="post">
                <tr>
                    <td><?= $contact['name'] ?></td>
                    <td><?= $contact['phone'] ?></td>
                    <td><?= $contact['email'] ?></td>
                    <td><input type="submit" value="Удалить"/></td>
                    <!-- "crutch" to transfer id to handler-->
                    <span style="display:none"><input type="" name="id" value="<?= $contact['id'] ?>"/></span>
                </tr>
                </form>
                <?php endforeach ?>
            </table>
            <!-- back button-->
            <a href="/"><input type="submit" value="Назад" /></a>
        </div>
        <?php include('views/footer.html'); ?>
    </body>
</html>
