-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 20, 2021 at 10:52 AM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 8.0.7

SET
SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET
time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `todos`
--

-- --------------------------------------------------------

--
-- Table structure for table `todos`
--

CREATE TABLE `todos`
(
    `id`          varchar(255) NOT NULL,
    `task`        varchar(100) NOT NULL,
    `starting_at` varchar(100) NOT NULL,
    `ends_at`     varchar(100) NOT NULL,
    `duration`    float        NOT NULL,
    `is_expired`  tinyint(1) NOT NULL,
    `done`        tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `todos`
--

INSERT INTO `todos` (`id`, `task`, `starting_at`, `ends_at`, `duration`, `is_expired`, `done`)
VALUES ('921f1b5f-e95f-4331-8caa-62d2cc6a9613', 'belajar kubernetess', '20 Aug 21 05:10 WIB', '23 Aug 21 04:10 WIB',
        4260, 0, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `todos`
--
ALTER TABLE `todos`
    ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;