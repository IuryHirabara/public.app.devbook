$('#new-publication').submit(createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.dislike-publication', dislikePublication);

$('#updatePublication').on('click', updatePublication);

$('.deletePublication').on('click', deletePublication);

function createPublication(e) {
    e.preventDefault();

    $.ajax({
        type: "POST",
        url: "/publications",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function () {
        window.location = '/home';
    })
        .fail(function (error) {
            Swal.fire(
                'Erro',
                'Não foi possível criar a publicação',
                'error'
            );
        });
}

function likePublication(e) {
    e.preventDefault();

    const element = $(e.target);
    const publicationId = element.closest('div').data('publication-id');

    element.prop('disabled', true);

    $.ajax({
        type: "POST",
        url: `publications/${publicationId}/like`
    })
        .done(function () {
            const likesCount = element.next('span');
            const likes = parseInt(likesCount.text());

            likesCount.text(likes + 1);

            element.addClass('dislike-publication');
            element.addClass('text-danger');
            element.removeClass('like-publication');
        })
        .fail(function () {
            Swal.fire(
                'Erro',
                'Error ao tentar curtir publicação',
                'error'
            );
        })
        .always(function () {
            element.prop('disabled', false);
        });
}

function dislikePublication(e) {
    e.preventDefault();

    const element = $(e.target);
    const publicationId = element.closest('div').data('publication-id');

    element.prop('disabled', true);

    $.ajax({
        type: "POST",
        url: `publications/${publicationId}/dislike`
    })
        .done(function () {
            const likesCount = element.next('span');
            const likes = parseInt(likesCount.text());

            likesCount.text(likes - 1);

            element.addClass('like-publication');
            element.removeClass('text-danger');
            element.removeClass('dislike-publication');
        })
        .fail(function () {
            Swal.fire(
                'Erro',
                'Error ao tentar descurtir publicação',
                'error'
            );
        })
        .always(function () {
            element.prop('disabled', false);
        });
}

function updatePublication(e) {
    $(this).prop('disabled', true);

    const publicationId = $(this).data('publication-id');

    $.ajax({
        type: "PUT",
        url: `/publications/${publicationId}`,
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    })
        .done(function () {
            Swal.fire(
                'Sucesso!',
                'Publicação editada com sucesso!',
                'success'
            ).then(function () {
                window.location = '/home'
            });
        })
        .fail(function () {
            Swal.fire(
                'Erro',
                'Error ao tentar editar publicação',
                'error'
            );
        })
        .always(function () {
            $('#updatePublication').prop('disabled', false);
        });
}

function deletePublication(e) {
    Swal.fire({
        title: 'Atenção!',
        text: 'Tem certeza que gostaria de excluir a publicação?',
        showCancelButton: true,
        cancelButtonText: 'Cancelar',
        icon: 'warning'
    }).then(function (confirmation) {
        if (!confirmation.value) return;

        const element = $(e.target);
        const publication = element.closest('div');
        const publicationId = publication.data('publication-id');

        element.prop('disabled', true);

        $.ajax({
            type: "DELETE",
            url: `/publications/${publicationId}`,
        })
            .done(function () {
                publication.fadeOut('slow', function () {
                    $(this).remove();
                });
            })
            .fail(function () {
                Swal.fire(
                    'Erro',
                    'Erro ao tentar excluir publicação',
                    'error'
                );
            });
    });
}