import { NextRequest, NextResponse } from 'next/server';
import getSession from './app/lib/session';

interface Routes {
  [key: string]: boolean;
}

const publicOnlyUrls: Routes = {
  '/': true,
  '/login': true,
  '/signup': true,
};

export async function middleware(request: NextRequest) {
  const session = await getSession();
  const exists = publicOnlyUrls[request.nextUrl.pathname];
  if (!session.id) {
    if (!exists) {
      console.log('exists', request.url);
      return NextResponse.redirect(new URL('/login', request.url));
    }
  } else {
    if (exists) {
      console.log('exists', request.url);
      return NextResponse.redirect(new URL('/profile', request.url));
    }
  }
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
