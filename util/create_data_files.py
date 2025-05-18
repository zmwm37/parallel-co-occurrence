import json
import pandas as pd
import sys

'''
create_data_files.py

Sample mpst dataset and create small, medium, and big files.
''' 

if __name__ == "__main__":
    p = "../data/raw/mpst_full_data.csv"
    SEED = 37
    if len(sys.argv) != 4:
        print("ERROR - three integers must be passed as arguments as small, medium, and large sample sizes.")
    else:
        small_n, medium_n, big_n = sys.argv[1], sys.argv[2], sys.argv[3]
        mpst = pd.read_csv(p)
        size_dict = {
            "small": small_n, 
            "medium": medium_n,
            "big": big_n
    }

        out_path = "../data/"
        for key, n in size_dict.items():
            mpst_sample = mpst.sample(int(n), random_state=SEED)
            sample_dict = mpst_sample[["imdb_id", "title", "plot_synopsis"]] \
                .rename(mapper = {"imdb_id": "Id", "title":"Title", "plot_synopsis":"Text"}, 
                        axis="columns") \
                .to_dict(orient="records")
            file = open(out_path+key+"/"+key+".txt" ,"w")
            for i in sample_dict:
                file.write(json.dumps(i)+"\n")
            file.close()