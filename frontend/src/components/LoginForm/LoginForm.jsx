import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import loginImage from '../../assets/logo.png';
import './LoginForm.css';

const LoginForm = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const response = await fetch('http://localhost:8080/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        // Cek apakah ini login pertama kali
        const isFirstLogin = !localStorage.getItem('token');

        if (isFirstLogin) {
          localStorage.setItem('token', 'dibimbing'); // Simpan token di localStorage
          setEmail('');
          setPassword('');
          navigate('/home');
        } else {
          const data = await response.json();
          console.log('Token:', data.token);

          setEmail('');
          setPassword('');

          navigate('/home');
        }
      } else {
        console.error('Login failed');
      }
    } catch (error) {
      console.error('Error during login:', error);
    }
  };

  return (
    <div className='bg-image'>
      <div className="login-container">
        <img src={loginImage} alt="Logo" style={{ width: '100px', height: '50px', display: 'block', margin: 'auto' }} />
        <form>
          <label>
            Email:
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="Enter your email"
            />
          </label>
          <br />
          <label>
            Password:
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Enter your password"
            />
          </label>
          <br />
          <button type="button" onClick={handleLogin}>
            LOGIN
          </button>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
