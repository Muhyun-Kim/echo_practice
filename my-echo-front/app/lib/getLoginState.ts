'use server';

import getSession from './session';

export async function getLoginState() {
  const session = await getSession();
  return !!session.id;
}
