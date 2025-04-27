import React from 'react';
import config from '../config';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const handlesignin = async (e: React.FormEvent) => {
        e.preventDefault();
        const response = await fetch(`${config.API_BASE_URL}/api/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email,
                password,
            }),
        });
        const data = await response.json();
        if (response.ok) {
            alert('Login successful!');
            
        } else {
            console.log(data);
            alert(`Login failed: ${data.message}. Please register to use the application.`);
        }
    }
    return (
        <form onSubmit={handlesignin}>
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
            <input type="email" className="form-control" placeholder="name@example.com" required
                onChange={e => setEmail(e.target.value)} />
            <input type="password" className="form-control" placeholder="Password" required
                onChange={e => setPassword(e.target.value)} />
            <button className="btn btn-primary w-100 py-2" type="submit">Sign in</button>
        </form>
    );
};

export default Login;