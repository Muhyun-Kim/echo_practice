export async function apiFetch(url: string, options: RequestInit = {}) {
  options.headers = {
    'Content-Type': 'application/json',
    ...options.headers,
  };

  options.credentials = 'include';

  const response = await fetch(url, options);
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || 'API request failed');
  }

  return response.json();
}
