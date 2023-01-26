[![Wizdwarfs](https://img.shields.io/badge/Wizdwarfs-designed%20by-green)]()
[![Docker](https://github.com/ali2210/WizDwarf/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/ali2210/WizDwarf/actions/workflows/docker-publish.yml)
[![Go](https://github.com/ali2210/WizDwarf/actions/workflows/go.yml/badge.svg)](https://github.com/ali2210/WizDwarf/actions/workflows/go.yml)         
[![OSSAR](https://github.com/ali2210/WizDwarf/actions/workflows/ossar-analysis.yml/badge.svg)](https://github.com/ali2210/WizDwarf/actions/workflows/ossar-analysis.yml)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/4fa84cf43d69415aa8cb51cad7b73a4f)](https://www.codacy.com/gh/ali2210/WizDwarf/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ali2210/WizDwarf&amp;utm_campaign=Badge_Grade)
[![CircleCI](https://circleci.com/gh/ali2210/WizDwarf/tree/master.svg?style=svg)](https://circleci.com/gh/ali2210/WizDwarf/tree/master)
[![Reference](https://godoc.org/github.com/ali2210/WizDwarf?status.svg)](https://pkg.go.dev/github.com/ali2210/wizdwarf)
[![Gitter](https://badges.gitter.im/wizdwarfs/futuristic-tech-dev.svg)](https://gitter.im/wizdwarfs/futuristic-tech-dev?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![GitNFT](https://img.shields.io/badge/%F0%9F%94%AE-Open%20in%20GitNFT-darkviolet?style=social)](https://gitnft.quine.sh/app/commits/list/repo/WizDwarf)
[![ReportCard](https://goreportcard.com/badge/github.com/ali2210/WizDwarf)](https://goreportcard.com/report/github.com/ali2210/WizDwarf)

 
 # About Project

          ## Abstract
         
                  The purpose of this project is to monitor the changes that have been made during the life. These changes happen due to different factors i.e environmental changes , outbreak dieases infection, etc.

         ## Scope or boardway

                  1. Genome Sequencing Processing (Analysis your genome).
                  2. Climate Genome Processing (a model that will cross reference climate and genome).
                  3. iGenome simulation  ( visualize genome through data for better understanding of the structure).


Genome Sequencing processing Mind Maps  

1.  [Genome_Sequence](https://drive.google.com/file/d/1cZfKNakiSEao_1v0JNIhEHHA2WtjYHNe/view?usp=share_link)


## Download & Install

      For Docker

                  Version >= 20.10.12


      ** For Linux ** 
             
             Either developer will forked the project or directly download from the repository as source file .
             Please choose latest source version for your machine. Another option is to download from dockerHub.

      ** For Windows & Mac **

             Currently docker-image is available on dockerHub for your machine. 

## Command-Line-Interface

      Developer will have start either with docker-image or composer.

      ** Docker-Image **
            
                  Forked the Project in your machine ; change the directory 
                  
                  $ git clone https://github.com/ali2210/WizDwarf.git && cd WizDwarf/
                  
                  $ docker run --rm -hostname=wizdwarfs -p 127.0.0.1:5000:5000 -v /app/app_data --net=host -it  wizdwarfs/wizdwarfs-v0.0.2     

      ** Composer ** 
                  
                  $ wget https://github.com/ali2210/WizDwarf/blob/master/docker-compose.yml && docker compose up 


## For Enterprise


            In Enterprize Option we are offering enchance security mechanism to protect your keys from bad actors.
            In that case developer will follow following instructions.  

            ** Install Vault **

                     https://developer.hashicorp.com/vault/docs/install

                     $ vault version

                     Version >= v1.11.0 

            ** Running Vault Agent on your machine **


                     ctrl + shift + T (Open new terminal) 
       
                     $  vault server -config=vault/config.hcl 
                     ctrl + shift + T (Open new terminal)

                     $ export VAULT_ADDR='http://127.0.0.1:8200'

                     https://developer.hashicorp.com/vault/tutorials/getting-started/getting-started-deploy

                     Open new tab type https://127.0.0.1:8200/ui and paste your vault login

                     Click + Enable New Engine Option and Choose KV option.

                     Add Path name "appsecret" . Without inverted commas;
                     
                     Add "hello" as input in secret data field  And click save option 

                     # Project changes

                     Paste token in creds.hci (token) and save the project.

                     Open vault folder and create new directory name "data"

                     ctrl + shift + T (Open new terminal) 

                     $ docker build -t wizdwarfs/wizdwarfs-v0.0.2 .

                     wait for compilation and run docker image like this
                     
                     $ docker run --rm -hostname=wizdwarfs -p 127.0.0.1:5000:5000 -v /app/app_data --net=host -it  wisdomenigma/wizdwarfs-v0.0.2

                     Make sure Developer Mode is enabled, Open browser and type 127.0.0.1:5000/

                     Congratulations Decentralized running on your machine.
   
## Remove & rebuild 
              
              Stop the Vault Server 
              $ pgrep -f vault | xargs kill

              $ rm -r ./vault/data inside the project (cli)
              
              Open creds.hcl replace token like this 
              Token_Auth = " " and save the project

              $ docker build -t wisdomenigma/wizdwarfs-v0.0.2 . or re-download from dockerHub

## API

   ### [GET]
        
       "/home"                      Home Page
	 "/signup"                    Signup User
	 "/login"                     Login User
	 "/dashboard"                 User Dashboard
       "/dashboard/profile"         User Profile
       "/dashboard/profile/edit"    Update Your Profile
       "/dashboard/profile/view"    View Your Profile   
       "/logout"                    Logout the user account
       "/error"                     Server Error
       "/dashboard/dvault"          Keygen manager 
       "/phenylalanine"             Phenylalanine protein
       "/leucine"                   Leucine protein
       "/isoleucine"                Isoleucine protein
       "/methionine"                Methionine protein
       "/valine"                    Valine protein
       "/serine"                    Serine protein
       "/proline"                   Proline protein
       "/threonine"                 Threonine protein
       "/alanine"                   Alanine protein
       "/tyrosine"                  Tyrosine protein
       "/histidine"                 Histidine protein
       "/glutamine"                 Glutamine protein
       "/asparagine"                Asparagine protein
       "/lysine"                    Lysine protein
       "/aspartic"                  Aspartic protein
       "/glutamic"                  Glutamic protein
       "/cysteine"                  Cysteine protein
       "/tryptophan"                Tryptophan protein
       "/arginine"                  Arginine protein
       "/glycine"                   Glycine protein
       "/stop"                       Stop Codon
    
   ### [POST]
       
       "/home"                HOME Page
	 "/signup"              Signup user
	 "/login"               Login User
	 "/dashboard"           User Dashboard
       "/dashboard/profile"   User Profile
       "/logout"              Logout the user account


## Features

       Visualization Option (BAR & LINE CHART)
       Optimized and clean code
       Vault for entreprizes 
       Keygen wallet
       Profile creation
       Encrypted content events
       Genome Sequence Matching
       VR Supported

       
## Discusion
     
   Developers will start conversation over project on various platforms. All yours queries resolved in community chat room. In case if you have an error or any idea regarding project improvement, contact us @ wizdwarfs@gmail.com  
   
   https://app.slack.com/client/T02AQ62EHHR/C02ALUWQ4LV/93ae0ff54319133c57487e772c8e0f1045690945 <span>,</span>

   <object>
              <a href='https://app.element.io/#/room/!XdCqKpBpqSSgLLcNPI:matrix.org' target='blank'>Element<a> <span>,</span> </object>
   <object>            
              <a href='https://www.reddit.com/r/wizdwarfs/a'> wizdwarfs_reddit </a>
   </object>


## Legal Business Option

       A happy customer is more precious than Gold mine. Payment is accepted only in bitcoin.  Distributors will owe licences, brand trademarks, copyright through legal business option.


## Licencing 

       This project is licensed under the Mozilla Public License (http://www.mozilla.org/MPL/). Distributors and royallty, copyright and other rights are included in license file. This software owe by WisdomEnigma, Inc


## FAQ

  1. How charts is render on web-page ? 
      Refresh webpage after wait for few seconds (~3s). Because data is shared between servers on client request.

  2. "EMPTY OUTPUT "  will be printed on console when app initate ?
      Application process is not terminate completely, before re initate application wait for it to finish this task. Expected time 1 minute.

  3.  Application is not started with vault after remove configurations ?
       Make sure data directory exists in vault directory. Otherwise, create manually and re-compile the dockerfile.

  4. How to get vault token and where it has been saved to initate the app?
      Please follow up Vault-Deploy instructions. Once you will get token store in creds.hcl (token field) & re-compile the dockerfile.  

  5. Why wallet is not render any document ?
       Refresh webpage after wait for few seconds (~3s). Because data is shared between servers on client request. Once page refresh all data will prefectly render on your web page.