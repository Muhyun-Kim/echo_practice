import { getIronSession } from 'iron-session';
import { cookies } from 'next/headers';

interface SessionContent {
  id?: string;
}

export default function getSession() {
  return getIronSession<SessionContent>(cookies(), {
    cookieName: 'session-name',
    password: process.env.COOKIE_PASSWORD!,
  });
}
