import React, { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import config from '../config';

function useQuery() {
    return new URLSearchParams(useLocation().search);
}

const VerifyEmail = () => {
    const query = useQuery();
    const email = query.get('email') || '';
    const [code, setCode] = useState('');
    const navigate = useNavigate();
    const [message, setMessage] = useState('');

    useEffect(() => {
        const sendVerificationEmail = async () => {
            try {
                const response = await fetch(`${config.API_BASE_URL}/api/send-verification`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ email }),
                });
                const data = await response.json();
                if (response.ok) {
                    setMessage('Verification email sent! Please check your Gmail.');
                } else {
                    setMessage(data.error || 'Failed to send verification email.');
                }
            } catch (error) {
                console.error(error);
                setMessage('Error sending verification email.');
            }
        };

        if (email) {
            sendVerificationEmail();
        }
    }, [email]);

    const handleVerify = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const response = await fetch(`${config.API_BASE_URL}/api/verify`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, code }),
            });
            const data = await response.json();
            if (response.ok) {
                alert('Email verified successfully!');
                navigate('/login');
            } else {
                alert('Invalid verification code. Please try again.');
            }
        } catch (error) {
            console.error(error);
            alert('Error verifying email.');
        }
    };

    return (
        <div>
            <h2>Email Verification</h2>
            <p>{message}</p>
            <form onSubmit={handleVerify}>
                <input
                    type="text"
                    placeholder="Enter Verification Code"
                    value={code}
                    onChange={e => setCode(e.target.value)}
                    required
                />
                <button type="submit">Verify</button>
            </form>
        </div>
    );
};

export default VerifyEmail;
