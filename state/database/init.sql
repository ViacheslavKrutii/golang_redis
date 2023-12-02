-- Створення бази даних, якщо вона ще не існує
CREATE DATABASE IF NOT EXISTS inviteHistory;

-- Використання бази даних
USE inviteHistory;

-- Створення таблиці inviteHistory
CREATE TABLE IF NOT EXISTS inviteHistory (
    Player VARCHAR(255),
    History TEXT
);