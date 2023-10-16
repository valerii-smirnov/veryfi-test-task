package config

type StreamConfig struct {
	DocumentStreamName        string `env:"DOCUMENT_STREAM_NAME,required"`
	CreateDocumentSubjectName string `env:"CREATE_DOCUMENT_SUBJECT_NAME,required"`
	RemoveDocumentSubjectName string `env:"REMOVE_DOCUMENT_SUBJECT_NAME,required"`
}
