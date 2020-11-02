CREATE TABLE `room` (
  `kdno` int(64) NOT NULL,
  `kcno` int(64) NOT NULL,
  `ccno` int(64) NOT NULL,
  `kdname` varchar(1024) NOT NULL,
  `time` datetime NOT NULL,
  `papername` varchar(1024) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
