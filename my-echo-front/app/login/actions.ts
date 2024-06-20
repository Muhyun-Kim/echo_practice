'use server';

export async function login(prevState: any, formData: FormData) {
  const data = {
    email: formData.get('email'),
    password: formData.get('password'),
  };
  const loginUrl = `${process.env.NEXT_PUBLIC_API_URL}user/login`;
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
    console.log('Login successful:', result);
  } else {
    console.log('Login failed:', result);
  }
}
