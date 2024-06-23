'use server';

import getSession from '../lib/session';
import { redirect } from 'next/navigation';

export async function login(prevState: any, formData: FormData) {
  const data = {
    email: formData.get('email'),
    password: formData.get('password'),
  };
  const loginUrl = `${process.env.NEXT_PUBLIC_API_URL}/user/login`;
  const res = await fetch(loginUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(data),
  });

  const result = await res.json();
  if (res.ok) {
    const setCookieHeader = res.headers.get('set-cookie')!;

    const sessionNameValue = setCookieHeader.split(';')[0];
    const [sessionName, sessionValue] = sessionNameValue.split('=');

    const session = await getSession();
    session.id = sessionValue;
    await session.save();
    redirect('/profile');
  } else {
    console.log('Login failed:', result);
  }
  return result;
}
