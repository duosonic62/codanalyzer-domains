package factory

import "github.com/duosonic62/codanalyzer-domains/internal/code"

//LoadCodes はコードを指定されたファイルのパスから生成する
func LoadCodes(filePath string) (*[]code.StructureBase, error) {
	codeInputs, err := readCodesFromJson(filePath)
	if err != nil {
		return nil, err
	}

	codes, err := makeCodeStructureBase(codeInputs)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

//LoadCodes はトライアドコードを指定されたファイルのパスから生成する
func LoadTriads(filePath string) (*[]code.TriadStructureBase, error) {
	codeInputs, err := readCodesFromJson(filePath)
	if err != nil {
		return nil, err
	}

	triads, err := makeTriadStructureBase(codeInputs)
	if err != nil {
		return nil, err
	}

	return triads, nil
}