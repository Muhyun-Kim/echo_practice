import { NextRequest, NextResponse } from 'next/server';

interface Routes {
  [key: string]: boolean;
}

const publicOnlyUrls: Routes = {
  '/': true,
  '/login': true,
  '/signup': true,
};

export async function middleware(request: NextRequest) {
  const sessionCookie = request.cookies.get('session-name');
  const exists = publicOnlyUrls[request.nextUrl.pathname];

  if (!sessionCookie) {
    if (!exists) {
      return NextResponse.redirect(new URL('/login', request.url));
    }
  } else {
    if (exists) {
      return NextResponse.redirect(new URL('/home', request.url));
    }
  }
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
