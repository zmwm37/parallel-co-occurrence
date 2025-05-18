import json
import pandas as pd
import sys


if __name__ == '__main__':

    n = sys.argv[1]
    p = '../data/raw/mpst_full_data.csv'
    mpst = pd.read_csv(p)

    text_dict = mpst[['plot_synopsis']] \
        .rename(mapper = {'plot_synopsis':'Text'}, 
                axis='columns') \
        .to_dict(orient='records')
    
    vocab_dict = {}
    for i in text_dict:
        # parse string
        text = i['Text']
        text = text.lower().split()
        # iterate through string and check dict
        for s in text:
            if vocab_dict.get(s):
                vocab_dict[s] += 1
            else:
                vocab_dict[s] = 1

    sorted_vocab = sorted(vocab_dict.items(), key=lambda x:x[1])
    sorted_vocab.reverse()
    top_n_vocab = sorted_vocab[:int(n)]

    top_n_dict = {}
    for i, t in enumerate(top_n_vocab):
        word, _  = t
        top_n_dict[word] = i
    
    with open("vocab_map_"+n+".txt", "w") as outfile:
        json.dump(top_n_dict, outfile)
