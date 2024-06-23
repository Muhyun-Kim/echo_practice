'use client';

import { AppBar, Box, Button, Toolbar, Typography } from '@mui/material';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import { getLoginState } from '../lib/getLoginState';

export default function Navbar() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  useEffect(() => {
    async function loginState() {
      const session = await getLoginState();
      setIsLoggedIn(session);
    }
    loginState();
  }, []);

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
              <Link href="/logout" passHref>
                <Button color="inherit">logout</Button>
              </Link>
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
