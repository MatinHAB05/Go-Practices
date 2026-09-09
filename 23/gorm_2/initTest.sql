USE
    bankdb;
DROP TABLE IF EXISTS
    customer_accounts;
DROP TABLE IF EXISTS
    customer_loans;
DROP TABLE IF EXISTS
    loans;
DROP TABLE IF EXISTS
    customers;
DROP TABLE IF EXISTS
    accounts;
DROP TABLE IF EXISTS
    loans;
DROP TABLE IF EXISTS
    branches;
DROP TABLE IF EXISTS
    banks;




INSERT INTO banks (code,name,address) VALUES
('B001','Melli','Tehran'),
('B002','Saderat','Tehran'),
('B003','Mellat','Tehran'),
('B004','Tejarat','Tehran'),
('B005','Parsian','Tehran');
INSERT INTO branches (branch_bank_code,branch_number,address) VALUES
('B001','101','Tehran Tajrish'),
('B001','102','Tehran Vanak'),
('B001','103','Tehran Shariati'),

('B002','201','Shiraz Center'),
('B002','202','Shiraz Zand'),
('B002','203','Isfahan Chaharbagh'),

('B003','301','Tehran Azadi'),
('B003','302','Tehran Enghelab'),
('B003','303','Karaj Gohardasht'),

('B004','401','Tabriz Imam'),
('B004','402','Tabriz Valiasr'),
('B004','403','Mashhad Sajjad'),

('B005','501','Ahvaz Center'),
('B005','502','Ahvaz Kianpars'),
('B005','503','Qom Haram');
INSERT INTO accounts (account_number,acc_bank_code,acc_branch_number,balance,type) VALUES
('A1001','B001','101',500000,'pro'),
('A1002','B001','101',750000,'pro-max'),
('A1003','B001','102',200000,'trial'),
('A1004','B001','102',1200000,'pro-max'),
('A1005','B001','103',950000,'pro'),

('A2001','B002','201',300000,'limited'),
('A2002','B002','201',840000,'pro'),
('A2003','B002','202',1500000,'pro-max'),
('A2004','B002','203',420000,'trial'),

('A3001','B003','301',980000,'pro'),
('A3002','B003','302',670000,'pro'),
('A3003','B003','303',210000,'limited'),

('A4001','B004','401',1100000,'pro-max'),
('A4002','B004','402',330000,'trial'),
('A4003','B004','403',450000,'pro'),

('A5001','B005','501',600000,'pro'),
('A5002','B005','502',730000,'pro-max'),
('A5003','B005','503',280000,'trial');
INSERT INTO loans (loan_number,loan_bank_code,loan_branch_number,amount,type) VALUES
('L1001','B001','101',10000000,'business'),
('L1002','B001','102',6000000,'government'),
('L1003','B001','103',8000000,'business'),

('L2001','B002','201',5000000,'government'),
('L2002','B002','202',7200000,'business'),
('L2003','B002','203',3100000,'political'),

('L3001','B003','301',4500000,'business'),
('L3002','B003','302',9100000,'government'),

('L4001','B004','401',6200000,'political'),
('L4002','B004','403',7400000,'business'),

('L5001','B005','501',3800000,'government'),
('L5002','B005','503',5600000,'business');
INSERT INTO customers (ssn,name,address,phone) VALUES
('1000000001','Ali Ahmadi','Tehran','09120000001'),
('1000000002','Sara Mohammadi','Shiraz','09120000002'),
('1000000003','Reza Karimi','Tabriz','09120000003'),
('1000000004','Neda Hosseini','Tehran','09120000004'),
('1000000005','Hossein Rahimi','Karaj','09120000005'),
('1000000006','Maryam Kiani','Ahvaz','09120000006'),
('1000000007','Amir Jalali','Mashhad','09120000007'),
('1000000008','Zahra Moradi','Qom','09120000008'),
('1000000009','Sina Ghaffari','Isfahan','09120000009'),
('1000000010','Leila Sadeghi','Tehran','09120000010'),
('1000000011','Hamed Nouri','Tabriz','09120000011'),
('1000000012','Fatemeh Abbasi','Shiraz','09120000012');
INSERT INTO customer_accounts (ca_bank_code,ca_branch_number,ca_account_number,ca_customer_ssn) VALUES
('B001','101','A1001','1000000001'),
('B001','101','A1002','1000000002'),
('B001','102','A1003','1000000003'),
('B001','102','A1004','1000000004'),
('B001','103','A1005','1000000005'),

('B002','201','A2001','1000000006'),
('B002','201','A2002','1000000007'),
('B002','202','A2003','1000000008'),
('B002','203','A2004','1000000009'),

('B003','301','A3001','1000000010'),
('B003','302','A3002','1000000011'),
('B003','303','A3003','1000000012'),

('B004','401','A4001','1000000001'),
('B004','402','A4002','1000000002'),
('B004','403','A4003','1000000003'),

('B005','501','A5001','1000000004'),
('B005','502','A5002','1000000005'),
('B005','503','A5003','1000000006');
INSERT INTO customer_loans (cl_bank_code,cl_branch_number,cl_loan_number,cl_customer_ssn) VALUES
('B001','101','L1001','1000000001'),
('B001','102','L1002','1000000003'),
('B001','103','L1003','1000000005'),

('B002','201','L2001','1000000007'),
('B002','202','L2002','1000000008'),
('B002','203','L2003','1000000009'),

('B003','301','L3001','1000000010'),
('B003','302','L3002','1000000011'),

('B004','401','L4001','1000000002'),
('B004','403','L4002','1000000003'),

('B005','501','L5001','1000000004'),
('B005','503','L5002','1000000006');
