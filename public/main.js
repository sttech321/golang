document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('user-form');
    const userList = document.getElementById('user-list');

    const apiUrl = 'http://localhost:8080/users';

    const fetchUsers = async () => {
        const response = await fetch(apiUrl);
        const users = await response.json();
        userList.innerHTML = '';
        for (const id in users) {
            const user = users[id];
            const li = document.createElement('li');
            li.innerHTML = `
                ${user.name} (${user.email})
                <button onclick="editUser(${user.id})">Edit</button>
                <button onclick="deleteUser(${user.id})">Delete</button>
            `;
            userList.appendChild(li);
        }
    };

    const createUser = async (user) => {
        await fetch(apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(user),
        });
    };

    const updateUser = async (id, user) => {
        await fetch(`${apiUrl}/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(user),
        });
    };

    const deleteUser = async (id) => {
        await fetch(`${apiUrl}/${id}`, {
            method: 'DELETE',
        });
        fetchUsers();
    };

    const editUser = async (id) => {
        const response = await fetch(`${apiUrl}/${id}`);
        const user = await response.json();
        document.getElementById('user-id').value = user.id;
        document.getElementById('name').value = user.name;
        document.getElementById('email').value = user.email;
    };

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const id = document.getElementById('user-id').value;
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const user = { name, email };
        if (id) {
            await updateUser(id, user);
        } else {
            await createUser(user);
        }
        form.reset();
        fetchUsers();
    });

    fetchUsers();

    window.editUser = editUser;
    window.deleteUser = deleteUser;
});
