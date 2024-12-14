-- Description: This script is used to create the User table in the database
-- 1. User Table
CREATE TABLE UserProfile (
    user_id BIGSERIAL,
    username VARCHAR(50),
    password TEXT,
    salt TEXT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(255),
    active BOOLEAN DEFAULT TRUE,
    status VARCHAR(10) NOT NULL,
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- 2. Role Table
CREATE TABLE Role (
    role_id INT,
    type VARCHAR(50),
    name VARCHAR(100),
    active BOOLEAN DEFAULT TRUE,
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- 3. User_Role_Mapping Table
CREATE TABLE User_Role_Mapping (
    user_role_mapping_id INT,
    user_id BIGSERIAL,
    role_id INT,
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- 4. Permission Table
CREATE TABLE Permission (
    permission_id INT,
    type VARCHAR(50),
    name VARCHAR(100),
    active BOOLEAN DEFAULT TRUE,
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- 5. Role_Permission_Mapping Table
CREATE TABLE Role_Permission_Mapping (
    role_permission_mapping_id INT,
    role_id INT,
    permission_id INT,
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- 6. Resource Table
CREATE TABLE Resource (
    resource_id INT,
    name VARCHAR(100),
    type VARCHAR(50),
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- 7. Permission_Resource_Mapping Table
CREATE TABLE Permission_Resource_Mapping (
    permission_resource_mapping_id INT,
    permission_id INT,
    resource_id INT,
    created_by INT,
    created_date TIMESTAMP,
    updated_by INT,
    updated_date TIMESTAMP
);

-- Add Constraints

-- User Table Constraints
ALTER TABLE UserProfile ADD CONSTRAINT USER_ID_PK PRIMARY KEY (user_id);
ALTER TABLE UserProfile ADD CONSTRAINT USER_USERNAME_UQ UNIQUE (username);
ALTER TABLE UserProfile ADD CONSTRAINT USER_PASSWORD_UQ UNIQUE (password);
ALTER TABLE UserProfile ADD CONSTRAINT USER_SALT_UQ UNIQUE (salt);
ALTER TABLE UserProfile ADD CONSTRAINT USER_STATUS_NN CHECK (status IN ('ACTIVE', 'INACTIVE'));

-- Role Table Constraints
ALTER TABLE Role ADD CONSTRAINT ROLE_ID_PK PRIMARY KEY (role_id);
ALTER TABLE Role ADD CONSTRAINT ROLE_NAME_UQ_NN UNIQUE (name);


-- User_Role_Mapping Table Constraints
ALTER TABLE User_Role_Mapping ADD CONSTRAINT USER_ROLE_MAPPING_PK PRIMARY KEY (user_role_mapping_id);
ALTER TABLE User_Role_Mapping ADD CONSTRAINT USER_ROLE_MAPPING_USERID_FK FOREIGN KEY (user_id) REFERENCES UserProfile (user_id);
ALTER TABLE User_Role_Mapping ADD CONSTRAINT USER_ROLE_MAPPING_ROLEID_FK FOREIGN KEY (role_id) REFERENCES Role (role_id);

-- Permission Table Constraints
ALTER TABLE Permission ADD CONSTRAINT PERMISSION_ID_PK PRIMARY KEY (permission_id);
ALTER TABLE Permission ADD CONSTRAINT PERMISSION_NAME_UQ_NN UNIQUE (name);

-- Role_Permission_Mapping Table Constraints
ALTER TABLE Role_Permission_Mapping ADD CONSTRAINT ROLE_PERMISSION_MAPPING_ROLEID_FK FOREIGN KEY (role_id) REFERENCES Role (role_id);
ALTER TABLE Role_Permission_Mapping ADD CONSTRAINT ROLE_PERMISSION_MAPPING_PERMISSIONID_FK FOREIGN KEY (permission_id) REFERENCES Permission (permission_id);

-- Resource Table Constraints
ALTER TABLE Resource ADD CONSTRAINT RESOURCE_ID_PK PRIMARY KEY (resource_id);
ALTER TABLE Resource ADD CONSTRAINT RESOURCE_NAME_UQ_NN UNIQUE (name);

-- Permission_Resource_Mapping Table Constraints
ALTER TABLE Permission_Resource_Mapping ADD CONSTRAINT PERMISSION_RESOURCE_MAPPING_PERMISSION_ID_FK FOREIGN KEY (permission_id) REFERENCES Permission (permission_id);
ALTER TABLE Permission_Resource_Mapping ADD CONSTRAINT PERMISSION_RESOURCE_MAPPING_RESOURCE_ID_FK FOREIGN KEY (resource_id) REFERENCES Resource (resource_id);
