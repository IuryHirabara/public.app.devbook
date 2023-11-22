$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#delete-user').on('click', deleteUser);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);

function unfollow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: 'POST'
    }).done(function () {
        window.location = `/users/${userId}`;
    }).fail(function () {
        Swal.fire('Erro', 'Não foi possível deixar de seguir o usuário', 'error');
        $(this).prop('disabled', false);
    });
}

function follow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: 'POST'
    }).done(function () {
        window.location = `/users/${userId}`;
    }).fail(function () {
        Swal.fire('Erro', 'Não foi possível seguir o usuário', 'error');
        $(this).prop('disabled', false);
    });
}

function edit(e) {
    e.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: 'PUT',
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        }
    }).done(function () {
        Swal.fire(
            'Sucesso',
            'Usuário editado com sucesso',
            'success'
        ).then(function () {
            window.location = '/profile';
        });
    }).fail(function () {
        Swal.fire('Erro', 'Não foi possível editar o usuário', 'error');
    });
}

function updatePassword(e) {
    e.preventDefault();

    const newPassword = $('#new-password').val();
    const rePassword = $('#re-password').val();
    if (newPassword !== rePassword) {
        Swal.fire('Aviso', 'As senhas não conferem', 'warning');
        return;
    }

    const currentPassword = $('#current-password').val();

    $.ajax({
        type: "POST",
        url: "/update-password",
        data: {
            new: newPassword,
            current: currentPassword
        }
    }).done(function () {
        Swal.fire(
            'Sucesso',
            'Senha atualizada com sucesso',
            'success'
        ).then(function () {
            window.location = '/profile';
        });
    }).fail(function () {
        Swal.fire('Erro', 'Não foi possível atualizar a senha', 'error');
    });
}

function deleteUser() {
    Swal.fire({
        title: 'Atenção!',
        text: 'Tem certeza que gostaria de excluir a sua conta? Essa ação é irreversível',
        showCancelButton: true,
        cancelButtonText: 'Cancelar',
        icon: 'warning'
    }).then(function (confirmation) {
        console.log('entrou');
        if (!confirmation.value) return;
        console.log('confirmou');

        $.ajax({
            type: "DELETE",
            url: "/delete-user",
        }).done(function () {
            Swal.fire(
                'Sucesso',
                'Usuário exclúido com sucesso',
                'success'
            ).then(function () {
                window.location = '/logout';
            });
        }).fail(function () {
            Swal.fire('Erro', 'Não foi possível excluir o usuário', 'error');
        });
    });
}