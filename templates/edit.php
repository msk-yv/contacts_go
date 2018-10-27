<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf8">
        <title>Редактирование контактов</title>
        <link rel="stylesheet" href="style.css" />
    </head>
    <body>
        <?php include('views/header.html'); ?>
        <div class="wraperFirstTable">
            <table class="insert_search">
                <tr>
                    <th>Внести новый контакт</th>
                    <th> Поиск </th>
                </tr> 
                <tr>
                    <td><?php include('views/insertForm.php'); ?></td>
                    <td><?php include('views/searchForm.php'); ?></td>
                </tr>
            </table>
        </div>
        <h2>Редактирование контакта:</h2>
        <div class="wraperSecondTable">
            <table class="results">
                <tr>
                    <th>Имя контакта</th>
                    <th>Телефон</th>
                    <th>Электронная почта</th>
                    <th></th>
                </tr> 
                <?php foreach($contacts as $contact): ?>
                <form action="/models/edit.php" method="post">
                    <tr>
                        <td><input type="text" name='name' value="<?=$contact['name']?>" required/></td>
                        <td><input type="tel" name="phone" value="<?=$contact['phone']?>" /></td>
                        <td><input type="email" name="email" value="<?=$contact['email']?>" /></td>
                        <td><input type="submit" value="Изменить"/></td>
                        <!-- "crutch" to transfer id to handler-->
                        <span style="display:none"><input type="" name="id" value="<?=$contact['id']?>"/></span>
                    </tr>
                </form>
                <?php endforeach ?>
            </table>
        </div>
        <?php include('views/footer.html'); ?>
    </body>
</html>
