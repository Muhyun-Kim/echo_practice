'use client';

import { AppBar, Box, Button, Toolbar, Typography } from '@mui/material';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import { deleteSession, getSession } from '../lib/session';
import { useRouter } from 'next/navigation';

export default function Navbar() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const router = useRouter();
  useEffect(() => {
    async function loginState() {
      const session = await getSession();
      if (session) {
        setIsLoggedIn(true);
      } else {
        setIsLoggedIn(false);
      }
    }
    loginState();
  }, []);

  const handleLogout = async () => {
    await deleteSession();
    router.push('/login');
  };
  return (
    <AppBar position="static">
      <Toolbar sx={{ justifyContent: 'space-between' }}>
        <ul
          style={{ display: 'flex', listStyle: 'none', margin: 0, padding: 0 }}
        >
          <li>home</li>
        </ul>
        <ul
          style={{ display: 'flex', listStyle: 'none', margin: 0, padding: 0 }}
        >
          {isLoggedIn ? (
            <li>
              <Button color="inherit" type="button" onClick={handleLogout}>
                logout
              </Button>
            </li>
          ) : (
            // logout state
            <Link href="/login" passHref>
              <Button color="inherit">login</Button>
            </Link>
          )}
        </ul>
      </Toolbar>
    </AppBar>
  );
}
