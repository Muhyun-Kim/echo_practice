'use server';

import { cookies } from 'next/headers';

export async function getSession() {
  const cookie = cookies().get('session-name');
  return cookie;
}

export async function deleteSession() {
  cookies().delete('session-name');
}
