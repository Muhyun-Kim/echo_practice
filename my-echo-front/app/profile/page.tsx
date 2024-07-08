'use client';

import { useEffect, useState } from 'react';
import { User, UserProfile } from '../types/user';
import { get } from 'http';
import { fetchGetProfile } from './actions';

export default function Profile() {
  const [profile, setProfile] = useState<UserProfile | null>(null);
  useEffect(() => {
    async function getProfile() {
      const res = await fetchGetProfile();
      console.log(res);
      setProfile(res);
    }
    getProfile();
  }, []);
  return (
    <div>
      <h1>{profile?.email}</h1>
    </div>
  );
}
