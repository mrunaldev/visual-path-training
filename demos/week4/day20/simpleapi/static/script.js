// Function to load and display users
async function loadUsers() {
    try {
        const response = await fetch('/users');
        const users = await response.json();
        
        const usersList = document.getElementById('usersList');
        usersList.innerHTML = '';

        users.forEach(user => {
            const userCard = document.createElement('div');
            userCard.className = 'user-card';
            
            userCard.innerHTML = `
                <div class="user-info">
                    <strong>${user.name}</strong> (${user.age} years)
                </div>
                <div class="user-actions">
                    <button class="delete-btn" onclick="deleteUser(${user.id})">Delete</button>
                </div>
            `;
            
            usersList.appendChild(userCard);
        });
    } catch (error) {
        console.error('Error loading users:', error);
    }
}

// Function to create a new user
async function createUser(event) {
    event.preventDefault();
    
    const nameInput = document.getElementById('name');
    const ageInput = document.getElementById('age');
    
    const user = {
        name: nameInput.value,
        age: parseInt(ageInput.value)
    };
    
    try {
        const response = await fetch('/users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        });
        
        if (response.ok) {
            nameInput.value = '';
            ageInput.value = '';
            loadUsers();
        } else {
            console.error('Error creating user:', response.statusText);
        }
    } catch (error) {
        console.error('Error creating user:', error);
    }
}

// Function to delete a user
async function deleteUser(id) {
    try {
        const response = await fetch(`/users/${id}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            loadUsers();
        } else {
            console.error('Error deleting user:', response.statusText);
        }
    } catch (error) {
        console.error('Error deleting user:', error);
    }
}

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    loadUsers();
    document.getElementById('createUserForm').addEventListener('submit', createUser);
});
