$('#register-form').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    if ($('#password').val() !== $('#re-password').val()) {
        Swal.fire(
            'Aviso',
            'As senhas não coincidem!',
            'error'
        );
        return;
    }

    $.ajax({
        url: '/users',
        method: 'POST',
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            password: $('#password').val()
        }
    }).done(function () {
        Swal.fire(
            'Sucesso',
            'Usuário cadastrado com sucesso!',
            'success'
        ).then(function () {
            $.ajax({
                type: "POST",
                url: "/login",
                data: {
                    email: $('#email').val(),
                    password: $('#password').val()
                }
            }).done(function () {
                window.location = '/home';
            }).fail(function () {
                Swal.fire(
                    'Erro',
                    'Erro ao autenticar usuário',
                    'error'
                );
            });
        });
    }).fail(function () {
        Swal.fire(
            'Erro',
            'Erro ao cadastrar usuário',
            'error'
        );
    });
}