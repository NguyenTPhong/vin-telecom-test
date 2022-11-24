SELECT 'CREATE DATABASE telecom'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'telecom');\gexec


