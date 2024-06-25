'use client';

import { Box, Button, Container, TextField, Typography } from '@mui/material';
import React, { useState } from 'react';
import { apiFetch } from '../lib/api';
import { redirect } from 'next/dist/server/api-utils';
import { useRouter } from 'next/navigation';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const data = { email, password };
    const loginUrl = `${process.env.NEXT_PUBLIC_API_URL}/user/login`;
    try {
      const res = await apiFetch(loginUrl, {
        method: 'POST',
        body: JSON.stringify(data),
      });
      if (res.message === 'Login successful') {
        router.push('/profile');
      }
    } catch (err) {
      console.log(err);
    }
  };
  return (
    <Container component="main">
      <Box className="mt-8 flex flex-col items-center">
        <Typography component="h1" variant="h5">
          Log In
        </Typography>
        <Box component="form" onSubmit={handleSubmit}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="password"
            label="Password"
            name="password"
            autoComplete="current-password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button type="submit" fullWidth variant="contained">
            Log In
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
