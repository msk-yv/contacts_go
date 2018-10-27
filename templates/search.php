<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf8">
        <title>Список контактов</title>
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
    <h2>Search results:</h2>
    <div class="wraperSecondTable">
    <table class="results">
        <tr>
            <th>Имя контакта</th>
            <th>Телефон</th>
            <th>Электронная почта</th>
            <th></th>
            <th></th>
        </tr> 
        <?php foreach($contacts as $contact): ?>
        <tr>
            <td><?=$contact['name']?></td>
            <td><?=$contact['phone']?></td>
            <td><?=$contact['email']?></td>
            <td><a href="/edit.php?id=<?=$contact['id']?>">Редактировать</a></td>
            <td><a href="/delete.php?id=<?=$contact['id']?>">Удалить</a></td>
        </tr>
        <?php endforeach ?>
    </table>
    </div>
    <?php include('views/footer.html'); ?>
</body>
</html>
