package errors

const (
	//ErrorRequired campo obrigatorio
	ErrorRequired = "Obrigatório"
	//ErrorPasswordLength Tamanho da senha
	ErrorPasswordLength = "A senha deve ter entre 6 e 40 caracteres"
	//ErrorExisting campo unico ja cadastrado
	ErrorExisting = "Já cadastrado"
	//ErrorInvalidEmail formato do email
	ErrorInvalidEmail = "E-mail inválido"
	//ErrorInvalidCpf formato do cpf
	ErrorInvalidCpf = "CPF inválido"
	//ErrorInvalidCnpj formato do cnpj
	ErrorInvalidCnpj = "CNPJ inválido"
	//ErrorInvalidDoc formato do documento
	ErrorInvalidDoc = "Documento inválido"
	//ErrorInvalidCnh formato da CNH
	ErrorInvalidCnh = "CNH inválida"
	//ErrorInvalidRenavam formato do renavam
	ErrorInvalidRenavam = "Renavam inválido"
	//ErrorInvalidPlate formato da placa
	ErrorInvalidPlate = "Placa inválida"
	//ErrorInvalidDate formato da data
	ErrorInvalidDate = "Data inválida"
	//ErrorEmptyResults nenhum registro
	ErrorEmptyResults = "Nenhum registro localizado"
	//SuccessDestroyRegister registro deletado
	SuccessDestroyRegister = "Registro deletado com sucesso"
	//ErrorDestroyRegister registro não deletado
	ErrorDestroyRegister = "Não foi possível deletar registro"
	//ErrorCreateRegister registro nao cadastrado
	ErrorCreateRegister = "Não foi possível cadastrar registro"
	//ErrorUpdateRegister registro nao atualizado
	ErrorUpdateRegister = "Não foi possível atualizar registro"
	//ErrorInFields erros nos campos
	ErrorInFields = "Ooops, verifique os campos"
	//ErrorInAutenticate authenticate error
	ErrorInAutenticate = "Ooops, dados informados não conferem"
	//SuccessInAutenticate messsage success
	SuccessInAutenticate = "Autenticação efetuada com sucesso"
)
