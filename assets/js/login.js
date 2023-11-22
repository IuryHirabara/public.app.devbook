$('#login').on('submit', login);

function login(event) {
    event.preventDefault();

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
            'O e-mail ou senha fornecidos são inválidos',
            'error'
        );
    });
}