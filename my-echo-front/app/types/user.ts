export interface UserSession {
  id: number;
  email: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
  password: string | null;
  createdAt: string | null;
}

export interface UserProfile {
  id: number;
  email: string;
  username: string;
}
