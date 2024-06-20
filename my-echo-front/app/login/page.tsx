'use client';

import { Box, Button, Container, TextField, Typography } from '@mui/material';
import React from 'react';
import { useFormState } from 'react-dom';
import { login } from './actions';

export default function Login() {
  const [state, dispatch] = useFormState(login, null);
  return (
    <Container component="main">
      <Box className="mt-8 flex flex-col items-center">
        <Typography component="h1" variant="h5">
          Log In
        </Typography>
        <Box component="form" action={dispatch}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            id="password"
            label="Password"
            name="password"
            autoComplete="current-password"
          />
          <Button type="submit" fullWidth variant="contained">
            Log In
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
