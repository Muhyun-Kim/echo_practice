'use client';

import { AppBar, Box, Button, Toolbar, Typography } from '@mui/material';
import Link from 'next/link';

export default function Navbar() {
  return (
    <AppBar position="static">
      <Toolbar sx={{ justifyContent: 'space-between' }}>
        <ul
          style={{ display: 'flex', listStyle: 'none', margin: 0, padding: 0 }}
        >
          <li></li>
        </ul>
        <ul
          style={{ display: 'flex', listStyle: 'none', margin: 0, padding: 0 }}
        >
          <li>
            <Link href="/login" passHref>
              <Button color="inherit">login</Button>
            </Link>
          </li>
          <li>
            <Button color="inherit">logout</Button>
          </li>
        </ul>
      </Toolbar>
    </AppBar>
  );
}
