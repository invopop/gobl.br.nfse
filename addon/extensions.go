package addon

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
	"github.com/invopop/gobl/pkg/here"
)

// Brazilian extension keys required to issue NFS-e documents
const (
	ExtKeyCNAE            = "br-nfse-cnae"
	ExtKeyNBS             = "br-nfse-nbs"
	ExtKeyFiscalIncentive = "br-nfse-fiscal-incentive"
	ExtKeyISSLiability    = "br-nfse-iss-liability"
	ExtKeyISSImmunity     = "br-nfse-iss-immunity"
	ExtKeyService         = "br-nfse-service"
	ExtKeySimples         = "br-nfse-simples"
	ExtKeySimplesRegime   = "br-nfse-simples-regime"
	ExtKeySpecialRegime   = "br-nfse-special-regime"
	ExtKeyOperation       = "br-nfse-operation"
	ExtKeyTaxStatus       = "br-nfse-tax-status"
	ExtKeyTaxClass        = "br-nfse-tax-class"
	ExtKeySupplyMode      = "br-nfse-supply-mode"
	ExtKeyRelationship    = "br-nfse-relationship"
)

var extensions = []*cbc.Definition{
	{
		Key: ExtKeyCNAE,
		Name: i18n.String{
			i18n.EN: "CNAE code",
			i18n.PT: "Código CNAE",
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				The CNAE (National Classification of Economic Activities) code for a service.

				List of codes from the IBGE (Brazilian Institute of Geography and Statistics):

				* https://www.ibge.gov.br/en/statistics/technical-documents/statistical-lists-and-classifications/17245-national-classification-of-economic-activities.html
			`),
		},
		Pattern: `^\d{2}[\s\.\-\/]?\d{2}[\s\.\-\/]?\d[\s\.\-\/]?\d{2}$`,
	},
	{
		Key: ExtKeyNBS,
		Name: i18n.String{
			i18n.EN: "NBS code",
			i18n.PT: "Código NBS",
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				The NBS (Nomenclatura Brasileira de Serviços, Intangíveis e outras Operações que
				produzam Variações no Patrimônio) code for a service. NBS is the Brazilian
				classification system for services and intangibles, established by Decreto No.
				7.708/2012 and maintained by the MDIC (Ministério do Desenvolvimento, Indústria,
				Comércio e Serviços).

				Official list of codes from the MDIC:

				* https://www.gov.br/mdic/pt-br/assuntos/sdic/comercio-e-servicos/nbs-nomenclatura-brasileira-de-servicos
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Decree No. 7.708/2012 - Institution of the NBS and NEBS",
					i18n.PT: "Decreto Nº 7.708/2012 - Instituição da NBS e NEBS",
				},
				URL: "https://www.planalto.gov.br/ccivil_03/_ato2011-2014/2012/decreto/d7708.htm",
			},
			{
				Title: i18n.String{
					i18n.EN: "NBS - Brazilian Nomenclature of Services (MDIC)",
					i18n.PT: "NBS - Nomenclatura Brasileira de Serviços (MDIC)",
				},
				URL: "https://www.gov.br/mdic/pt-br/assuntos/sdic/comercio-e-servicos/nbs-nomenclatura-brasileira-de-servicos",
			},
		},
		Pattern: `^\d[\.\s]?\d{4}[\.\s]?\d{2}[\.\s]?\d{2}$`,
	},
	{
		Key: ExtKeyFiscalIncentive,
		Name: i18n.String{
			i18n.EN: "Fiscal Incentive",
			i18n.PT: "Incentivo Fiscal",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Has incentive",
					i18n.PT: "Possui incentivo",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Does not have incentive",
					i18n.PT: "Não possui incentivo",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates whether a party benefits from a fiscal incentive.
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "NFS-e ABRASF Taxpayer Guidance Manual (v2.04)",
					i18n.PT: "NFS-e ABRASF Manual de Orientação do Contribuinte (v2.04)",
				},
				URL: "https://abrasf.org.br/biblioteca/arquivos-publicos/nfs-e-manual-de-orientacao-do-contribuinte-2-04/download",
			},
		},
	},
	{
		Key: ExtKeyISSLiability,
		Name: i18n.String{
			i18n.EN: "ISS Liability",
			i18n.PT: "Exigibilidade ISS",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Liable",
					i18n.PT: "Exigível",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Not subject",
					i18n.PT: "Não incidência",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Exempt",
					i18n.PT: "Isenção",
				},
			},
			{
				Code: "4",
				Name: i18n.String{
					i18n.EN: "Export",
					i18n.PT: "Exportação",
				},
			},
			{
				Code: "5",
				Name: i18n.String{
					i18n.EN: "Immune",
					i18n.PT: "Imunidade",
				},
			},
			{
				Code: "6",
				Name: i18n.String{
					i18n.EN: "Suspended Judicially",
					i18n.PT: "Suspensa por Decisão Judicial",
				},
			},
			{
				Code: "7",
				Name: i18n.String{
					i18n.EN: "Suspended Administratively",
					i18n.PT: "Suspensa por Processo Administrativo",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates the ISS liability status, i.e., whether the ISS tax is due or not and why.

				List of codes from the national NFSe ABRASF (v2.04) model:

				* https://abrasf.org.br/biblioteca/arquivos-publicos/nfs-e-manual-de-orientacao-do-contribuinte-2-04/download
				(Section 10.2, Field B-38)
			`),
		},
	},
	{
		Key: ExtKeyISSImmunity,
		Name: i18n.String{
			i18n.EN: "ISS Immunity",
			i18n.PT: "Identificação da Imunidade do ISSQN",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Assets, income or services of government entities between each other (CF88, Art 150, VI, a)",
					i18n.PT: "Patrimônio, renda ou serviços, uns dos outros (CF88, Art 150, VI, a)",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Religious entities and temples of any faith, including their assistance and charitable organisations (CF88, Art 150, VI, b)",
					i18n.PT: "Entidades religiosas e templos de qualquer culto, inclusive suas organizações assistenciais e beneficentes (CF88, Art 150, VI, b)",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Assets, income or services of political parties, including their foundations, workers' unions, and non-profit education and social assistance institutions, as required by law (CF88, Art 150, VI, c)",
					i18n.PT: "Patrimônio, renda ou serviços dos partidos políticos, inclusive suas fundações, das entidades sindicais dos trabalhadores, das instituições de educação e de assistência social, sem fins lucrativos, atendidos os requisitos da lei (CF88, Art 150, VI, c)",
				},
			},
			{
				Code: "4",
				Name: i18n.String{
					i18n.EN: "Books, newspapers, periodicals and the paper used to print them (CF88, Art 150, VI, d)",
					i18n.PT: "Livros, jornais, periódicos e o papel destinado a sua impressão (CF88, Art 150, VI, d)",
				},
			},
			{
				Code: "5",
				Name: i18n.String{
					i18n.EN: "Music recordings and music videos produced in Brazil with musical or literary-musical works by Brazilian authors or works performed by Brazilian artists, and the physical media or digital files containing them, except at the industrial replication stage of laser-read optical media (CF88, Art 150, VI, e)",
					i18n.PT: "Fonogramas e videofonogramas musicais produzidos no Brasil contendo obras musicais ou literomusicais de autores brasileiros e/ou obras em geral interpretadas por artistas brasileiros bem como os suportes materiais ou arquivos digitais que os contenham, salvo na etapa de replicação industrial de mídias ópticas de leitura a laser. (CF88, Art 150, VI, e)",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				The constitutional immunity that applies when the ISS of a service is immune,
				that is, when the ISS tax combo has the ~br-nfse-iss-liability~ code ~5~.

				Maps to the ~tpImunidade~ field in the NFS-e national layout.
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Annex I - DPS and NFS-e Layout v1.01",
					i18n.PT: "Anexo I - Leiaute DPS e NFS-e v1.01",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/documentacao-atual/anexo_i-sefin_adn-dps_nfse-snnfse-v1-01-20260209.xlsx",
			},
		},
	},
	{
		Key: ExtKeyService,
		Name: i18n.String{
			i18n.EN: "Service Code",
			i18n.PT: "Código Item Lista Serviço",
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				The service code as defined by the municipality. Typically, one of the codes listed
				in the Lei Complementar 116/2003, but municipalities can make their own changes.

				For further details on the list of possible codes, see:

				* https://www.planalto.gov.br/ccivil_03/leis/lcp/lcp116.htm

				For suppliers issuing through the NFS-e Nacional, the 6-digit national service code
				instead, which maps to the ~cTribNac~ field in the national layout. List of possible
				values:

				* https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/documentacao-atual/anexo_b-nbs2-lista_servico_nacional-snnfse-v1-01-20260122.xlsx
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Annex B - National Service List and NBS v1.01",
					i18n.PT: "Anexo B - Lista de Serviço Nacional e NBS v1.01",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/documentacao-atual/anexo_b-nbs2-lista_servico_nacional-snnfse-v1-01-20260122.xlsx",
			},
		},
	},
	{
		Key: ExtKeySimples,
		Name: i18n.String{
			i18n.EN: "Opting for \"Simples Nacional\" regime",
			i18n.PT: "Optante pelo Simples Nacional",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Opt-in",
					i18n.PT: "Optante",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Opt-out",
					i18n.PT: "Não optante",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates whether a party is opting for the "Simples Nacional" (Regime Especial
				Unificado de Arrecadação de Tributos e Contribuições devidos pelas Microempresas e
				Empresas de Pequeno Porte) tax regime
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "NFS-e ABRASF Taxpayer Guidance Manual (v2.04)",
					i18n.PT: "NFS-e ABRASF Manual de Orientação do Contribuinte (v2.04)",
				},
				URL: "https://abrasf.org.br/biblioteca/arquivos-publicos/nfs-e-manual-de-orientacao-do-contribuinte-2-04/download",
			},
		},
	},
	{
		Key: ExtKeySimplesRegime,
		Name: i18n.String{
			i18n.EN: "Simples Nacional Tax Assessment Regime",
			i18n.PT: "Regime de Apuração Tributária pelo Simples Nacional",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Federal and municipal taxes through the Simples Nacional",
					i18n.PT: "Regime de apuração dos tributos federais e municipal pelo SN",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Federal taxes through the Simples Nacional, ISS under the municipal legislation",
					i18n.PT: "Regime de apuração dos tributos federais pelo SN e o ISSQN pela NFS-e conforme respectiva legislação municipal do tributo",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Federal and municipal taxes under their own legislation",
					i18n.PT: "Regime de apuração dos tributos federais e municipal pela NFS-e conforme respectivas legislações federal e municipal de cada tributo",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				How a supplier opting for the Simples Nacional as ME/EPP pays its federal taxes and
				ISS, which changes when it exceeds a Simples Nacional sublimit or limit. On the
				national system, required for ME/EPP suppliers and not allowed for any other.

				Maps to the ~regApTribSN~ field in the NFS-e national layout.
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Annex I - DPS and NFS-e Layout v1.01",
					i18n.PT: "Anexo I - Leiaute DPS e NFS-e v1.01",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/documentacao-atual/anexo_i-sefin_adn-dps_nfse-snnfse-v1-01-20260209.xlsx",
			},
		},
	},
	{
		Key: ExtKeySpecialRegime,
		Name: i18n.String{
			i18n.EN: "Special Tax Regime",
			i18n.PT: "Regime Especial de Tributação",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Municipal micro-enterprise",
					i18n.PT: "Microempresa municipal",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Estimated",
					i18n.PT: "Estimativa",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Professional Society",
					i18n.PT: "Sociedade de profissionais",
				},
			},
			{
				Code: "4",
				Name: i18n.String{
					i18n.EN: "Cooperative",
					i18n.PT: "Cooperativa",
				},
			},
			{
				Code: "5",
				Name: i18n.String{
					i18n.EN: "Single micro-entrepreneur (MEI)",
					i18n.PT: "Microempreendedor individual (MEI)",
				},
			},
			{
				Code: "6",
				Name: i18n.String{
					i18n.EN: "Micro-enterprise or Small Business (ME EPP)",
					i18n.PT: "Microempresa ou Empresa de Pequeno Porte (ME EPP).",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates a special tax regime that a party is subject to.
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "NFS-e ABRASF Taxpayer Guidance Manual (v2.04)",
					i18n.PT: "NFS-e ABRASF Manual de Orientação do Contribuinte (v2.04)",
				},
				URL: "https://abrasf.org.br/biblioteca/arquivos-publicos/nfs-e-manual-de-orientacao-do-contribuinte-2-04/download",
			},
		},
	},
	{
		Key: ExtKeyOperation,
		Name: i18n.String{
			i18n.EN: "Operation Indicator",
			i18n.PT: "Indicador da operação",
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates the operation type for the determination of the IBS and CBS taxes by the
				tax authorities.

				Maps to the ~cIndOp~ field in the NFS-e national layout.

				List of possible values:

				* https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/rtc/anexovii-indop_ibscbs_v1-00-00.xlsx
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Technical Note SE/CGNFS-e nº 004",
					i18n.PT: "Nota Técnica SE/CGNFS-e nº 004",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/rtc-producao-restrita-piloto/nt-004-se-cgnfse-novo-layout-rtc-v2-00-20251210.pdf",
			},
			{
				Title: i18n.String{
					i18n.EN: "Annex VII - Operation Indicators Table IBS/CBS v1.00",
					i18n.PT: "Anexo VII - Tabela de Indicadores de Operação IBS/CBS v1.00",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/rtc/anexovii-indop_ibscbs_v1-00-00.xlsx",
			},
		},
		Pattern: `^\d{6}$`,
	},
	{
		Key: ExtKeyTaxStatus,
		Name: i18n.String{
			i18n.EN: "Tax Status Code (CST)",
			i18n.PT: "Código de situação tributária (CST)",
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates the tax status of the operation for the determination of the IBS and CBS
				taxes by the tax authorities.

				Maps to the ~CST~ field in the NFS-e national layout.

				List of possible values:

				* https://dfe-portal.svrs.rs.gov.br/DFE/ClassificacaoTributaria
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Technical Report RT 2025.002 - CST and cClassTrib Tables IBS/CBS",
					i18n.PT: "Informe Técnico RT 2025.002 - Tabelas CST e cClassTrib IBS/CBS",
				},
				URL: "https://www.nfe.fazenda.gov.br/portal/exibirArquivo.aspx?conteudo=gya58CS0dHU=",
			},
		},
		Pattern: `^\d{3}$`,
	},
	{
		Key: ExtKeyTaxClass,
		Name: i18n.String{
			i18n.EN: "Tax Classification Code",
			i18n.PT: "Código de classificação tributária",
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Indicates the tax classification code for the determination of the IBS and CBS taxes
				by the tax authorities.

				Maps to the ~cClassTrib~ field in the NFS-e national layout.

				List of possible values:

				* https://dfe-portal.svrs.rs.gov.br/DFE/ClassificacaoTributaria
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Technical Report RT 2025.002 - CST and cClassTrib Tables IBS/CBS",
					i18n.PT: "Informe Técnico RT 2025.002 - Tabelas CST e cClassTrib IBS/CBS",
				},
				URL: "https://www.nfe.fazenda.gov.br/portal/exibirArquivo.aspx?conteudo=gya58CS0dHU=",
			},
		},
		Pattern: `^\d{6}$`,
	},
	{
		Key: ExtKeySupplyMode,
		Name: i18n.String{
			i18n.EN: "Supply Mode",
			i18n.PT: "Modo de Prestação",
		},
		Values: []*cbc.Definition{
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Cross-border",
					i18n.PT: "Transfronteiriço",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Consumption in Brazil",
					i18n.PT: "Consumo no Brasil",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Temporary movement of individuals",
					i18n.PT: "Movimento Temporário de Pessoas Físicas",
				},
			},
			{
				Code: "4",
				Name: i18n.String{
					i18n.EN: "Consumption abroad",
					i18n.PT: "Consumo no Exterior",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				How a service exported or imported was supplied.

				Maps to the ~mdPrestacao~ field in the NFS-e national layout.
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Annex I - DPS and NFS-e Layout v1.01",
					i18n.PT: "Anexo I - Leiaute DPS e NFS-e v1.01",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/documentacao-atual/anexo_i-sefin_adn-dps_nfse-snnfse-v1-01-20260209.xlsx",
			},
		},
	},
	{
		Key: ExtKeyRelationship,
		Name: i18n.String{
			i18n.EN: "Relationship Between Parties",
			i18n.PT: "Vínculo entre as partes no negócio",
		},
		Values: []*cbc.Definition{
			{
				Code: "0",
				Name: i18n.String{
					i18n.EN: "No relationship",
					i18n.PT: "Sem vínculo com o Tomador/Prestador",
				},
			},
			{
				Code: "1",
				Name: i18n.String{
					i18n.EN: "Controlled",
					i18n.PT: "Controlada",
				},
			},
			{
				Code: "2",
				Name: i18n.String{
					i18n.EN: "Controlling",
					i18n.PT: "Controladora",
				},
			},
			{
				Code: "3",
				Name: i18n.String{
					i18n.EN: "Affiliated",
					i18n.PT: "Coligada",
				},
			},
			{
				Code: "4",
				Name: i18n.String{
					i18n.EN: "Head office",
					i18n.PT: "Matriz",
				},
			},
			{
				Code: "5",
				Name: i18n.String{
					i18n.EN: "Branch",
					i18n.PT: "Filial ou sucursal",
				},
			},
			{
				Code: "6",
				Name: i18n.String{
					i18n.EN: "Other relationship",
					i18n.PT: "Outro vínculo",
				},
			},
		},
		Desc: i18n.String{
			i18n.EN: here.Doc(`
				Relationship between the supplier and the foreign customer, as defined in
				art. 23 of Law 9.430/1996.

				Maps to the ~vincPrest~ field in the NFS-e national layout.
			`),
		},
		Sources: []*cbc.Source{
			{
				Title: i18n.String{
					i18n.EN: "Annex I - DPS and NFS-e Layout v1.01",
					i18n.PT: "Anexo I - Leiaute DPS e NFS-e v1.01",
				},
				URL: "https://www.gov.br/nfse/pt-br/biblioteca/documentacao-tecnica/documentacao-atual/anexo_i-sefin_adn-dps_nfse-snnfse-v1-01-20260209.xlsx",
			},
		},
	},
}
