
/* Criando o banco de dados */
CREATE TABLE link (
	url text not null primary key,
	titulo text not null,
	tags text not null,
	data_criacao timestamp not null,
	privado bool);
